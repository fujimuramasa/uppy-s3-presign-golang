package presigner

import (

    "time"
    "net/http"
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type RequestBody struct {

    FileName string `json:"filename,omitempty"`
    ContentType string `json:"contentType,omitempty"`
}

func Sign(w http.ResponseWriter, r *http.Request) {

    var body RequestBody

    err := json.NewDecoder(r.Body).Decode(&body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String("ap-northeast-1"),
    })

    svc := s3.New(sess)

    req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
        Bucket: aws.String("yuusha2021"),
        Key: aws.String(body.FileName),
    })
    url, err := req.Presign(15 * time.Minute)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
    }

    ret := map[string]interface{}{
        "url": url,
        "method": "PUT",
        "fields": map[string]interface{}{},
    }
    byt, _ := json.Marshal(ret)

    w.Header().Add("content-type", "application/json")
    w.Write(byt)
}
