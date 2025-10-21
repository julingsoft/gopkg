package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/gogf/gf/v2/frame/g"
)

type Elasticsearch struct {
	client *elasticsearch.Client
}

var (
	instance *Elasticsearch
	once     sync.Once
)

// GetInstance 返回 Elasticsearch 的单例实例
func GetInstance(config elasticsearch.Config) (*Elasticsearch, error) {
	var err error
	once.Do(func() {
		client, e := elasticsearch.NewClient(config)
		if e != nil {
			err = e
			return
		}
		instance = &Elasticsearch{
			client: client,
		}
	})
	if err != nil {
		return nil, fmt.Errorf("error creating the client: %w", err)
	}
	return instance, nil
}

// Client 返回 Elasticsearch 客户端
func (e *Elasticsearch) Client() *elasticsearch.Client {
	return e.client
}

// CreateIndex 创建索引
func (e *Elasticsearch) CreateIndex(ctx context.Context, indexName string) error {
	res, err := e.client.Indices.Create(
		indexName,
		e.client.Indices.Create.WithContext(ctx),
	)
	if err != nil {
		return fmt.Errorf("error creating index: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from server: %s", res.String())
	}

	g.Log().Printf(ctx, "Index %s created successfully", indexName)

	return nil
}

// CreateDocument 创建（索引）文档
func (e *Elasticsearch) CreateDocument(ctx context.Context, indexName string, document []byte, id string) error {
	req := esapi.IndexRequest{
		Index:      indexName,
		DocumentID: id,
		Body:       bytes.NewReader(document),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("error indexing document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from server: %s", res.String())
	}

	g.Log().Printf(ctx, "Document %s created successfully", id)

	return nil
}

// ReadDocument 读取文档
func (e *Elasticsearch) ReadDocument(ctx context.Context, indexName, docID string) (*any, error) {
	req := esapi.GetRequest{
		Index:      indexName,
		DocumentID: docID,
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return nil, fmt.Errorf("error getting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from server: %s", res.String())
	}

	var result struct {
		Source any `json:"_source"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &result.Source, nil
}

// UpdateDocument 更新文档
func (e *Elasticsearch) UpdateDocument(ctx context.Context, indexName, docID string, updateData map[string]interface{}) error {
	body, err := json.Marshal(map[string]interface{}{"doc": updateData})
	if err != nil {
		return fmt.Errorf("error marshaling update data: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      indexName,
		DocumentID: docID,
		Body:       bytes.NewReader(body),
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("error updating document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from server: %s", res.String())
	}

	g.Log().Printf(ctx, "Document %s updated successfully", docID)

	return nil
}

// DeleteDocument 删除文档
func (e *Elasticsearch) DeleteDocument(ctx context.Context, indexName, docID string) error {
	req := esapi.DeleteRequest{
		Index:      indexName,
		DocumentID: docID,
		Refresh:    "true",
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return fmt.Errorf("error deleting document: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from server: %s", res.String())
	}

	g.Log().Printf(ctx, "Document %s deleted successfully", docID)

	return nil
}

// SearchDocuments 搜索文档
func (e *Elasticsearch) SearchDocuments(ctx context.Context, indexName, query string) ([]any, error) {
	var buf bytes.Buffer
	searchQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": query,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(searchQuery); err != nil {
		return nil, fmt.Errorf("error encoding query: %w", err)
	}

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  &buf,
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return nil, fmt.Errorf("error searching documents: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response from server: %s", res.String())
	}

	var result struct {
		Hits struct {
			Hits []struct {
				Source any `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error decoding search response: %w", err)
	}

	var documents []any
	for _, hit := range result.Hits.Hits {
		documents = append(documents, hit.Source)
	}
	return documents, nil
}
