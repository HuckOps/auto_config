package reader

import (
	"errors"
	"github.com/huckops/auto_config/encoder"
	"github.com/huckops/auto_config/encoder/json"
	"github.com/huckops/auto_config/source"
	"github.com/imdario/mergo"
	"time"
)

type jsonReader struct {
	opts Options
	json encoder.Encoder
}

func (j *jsonReader) Merge(changes ...*source.ChangeSet) (*source.ChangeSet, error) {
	merged := map[string]interface{}{}
	for _, change := range changes {
		codec, ok := j.opts.Encoding[change.Format]
		if !ok {
			return nil, errors.New("get encoder failed")
		}

		var data map[string]interface{}
		// 转为结构体
		if err := codec.Decode(change.Data, &data); err != nil {
			return nil, err
		}
		// 结构体合并
		if err := mergo.Map(&merged, data, mergo.WithOverride); err != nil {
			return nil, err
		}
	}
	// 序列化为json
	b, err := j.json.Encode(merged)
	//b, err := json2.Marshal(merged)
	if err != nil {
		panic(err)
		return nil, err
	}
	cs := &source.ChangeSet{
		Timestamp: time.Now(),
		Data:      b,
		Format:    j.json.Type(),
	}
	cs.CheckSum = cs.Sum()

	return cs, nil
}
func (j *jsonReader) Scan(data []byte, v interface{}) error {
	jsonEncoding := j.opts.Encoding["json"]
	return jsonEncoding.Decode(data, v)

}
func NewReader(opts ...Option) Reader {
	options := NewOptions(opts...)
	return &jsonReader{
		json: json.NewEncoder(),
		opts: options,
	}
}
