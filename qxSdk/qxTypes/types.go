// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3

package qxTypes

type AllowCreateAndDirectUpload struct {
	FileName         string   `json:"fileName"`
	BucketKey        string   `json:"bucketKey"`
	FolderId         uint32   `json:"folderId,optional"`
	Remark           string   `json:"remark,optional"`
	Cover            string   `json:"cover,optional"`
	AlternativeCover []string `json:"alternativeCover,optional"`
}

type AllowCreateModelSasFile struct {
	UploadType       int32    `json:"uploadType"`
	FileName         string   `json:"fileName"`
	FolderId         uint32   `json:"folderId,optional"`
	Remark           string   `json:"remark,optional"`
	Cover            string   `json:"cover"`
	AlternativeCover []string `json:"alternativeCover,optional"`
	FileSize         int64    `json:"fileSize,optional"`
	MimeType         string   `json:"mimeType,optional"`
	ExternalUrl      string   `json:"externalUrl,optional"`
	BucketKey        string   `json:"bucketKey,optional"`
}

type AllowCreateModelSasFolder struct {
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Cover       string `json:"cover,optional"`
	ArchiveType int32  `json:"archiveType,optional"`
	ParentId    uint32 `json:"parentId,optional"`
}

type AllowUpdateModelSasFile struct {
	Id               uint32   `json:"id"`
	FolderId         uint32   `json:"folderId,optional"`
	Name             string   `json:"name,optional"`
	Remark           string   `json:"remark,optional"`
	Cover            string   `json:"cover"`
	AlternativeCover []string `json:"alternativeCover"`
}

type AllowUpdateModelSasFolder struct {
	Id          uint32 `json:"id"`
	Name        string `json:"name,optional"`
	Alias       string `json:"alias,optional"`
	Cover       string `json:"cover,optional"`
	ArchiveType int32  `json:"archiveType,optional"`
	ParentId    uint32 `json:"parentId"`
}

type AllowUpdateStatusModelSasFile struct {
	Id     uint32 `json:"id"`
	Status int32  `json:"status"`
}

type AllowUpdateStatusModelSasFolder struct {
	Id     uint32 `json:"id"`
	Status int32  `json:"status"`
}

type ApiCaptchaGenerateReq struct {
	Key       string  `json:"key,optional"`
	DotCount  int32   `json:"dotCount,optional"`
	MaxSkew   float64 `json:"maxSkew,optional"`
	KeyLong   int32   `json:"keyLong,optional"`
	ImgWidth  int32   `json:"imgWidth,optional"`
	ImgHeight int32   `json:"imgHeight,optional"`
}

type ApiCaptchaGenerateResp struct {
	Code      int32                      `json:"code"`
	Msg       string                     `json:"msg"`
	Path      string                     `json:"path"`
	RequestID string                     `json:"requestId"`
	Data      ApiCaptchaGenerateRespData `json:"data"`
}

type ApiCaptchaGenerateRespData struct {
	Key     string `json:"key,omitempty"`
	Id      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Answer  string `json:"answer,omitempty"`
	Img     string `json:"img,omitempty"`
}

type ApiRefreshReq struct {
	AccessKey    string `json:"accessKey"`
	RefreshToken string `json:"refreshToken"`
}

type ApiRefreshResp struct {
	Code      int32           `json:"code"`
	Msg       string          `json:"msg"`
	Path      string          `json:"path"`
	RequestID string          `json:"requestId"`
	Data      SignResultModel `json:"data"`
}

type ApiSignReq struct {
	AccessKey    string `json:"accessKey"`
	AccessSecret string `json:"accessSecret"`
}

type ApiSignResp struct {
	Code      int32           `json:"code"`
	Msg       string          `json:"msg"`
	Path      string          `json:"path"`
	RequestID string          `json:"requestId"`
	Data      SignResultModel `json:"data"`
}

type ApiSmsSendReq struct {
	Key    string   `json:"key"`
	Mobile string   `json:"mobile"`
	Params []string `json:"params,optional"`
}

type ApiSmsSendResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type ApiSmsSendRespData struct {
	RequestID string `json:"requestId"`
}

type BehavioralVerificationInitReq struct {
	Key     string `json:"key"`
	Service string `json:"service"`
	Type    string `json:"type"`
}

type BehavioralVerificationInitResp struct {
	Code      int32                              `json:"code"`
	Msg       string                             `json:"msg"`
	Path      string                             `json:"path"`
	RequestID string                             `json:"requestId"`
	Data      BehavioralVerificationInitRespData `json:"data"`
}

type BehavioralVerificationInitRespData struct {
	Id      string `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
	Answer  string `json:"answer,omitempty"`
	Img     string `json:"img,omitempty"`
}

type BehavioralVerificationVerifyReq struct {
	Id         string `json:"id"`
	Service    string `json:"service"`
	Type       string `json:"type"`
	VerifyCode string `json:"verifyCode"`
}

type BehavioralVerificationVerifyResp struct {
	Code      int32                                `json:"code"`
	Msg       string                               `json:"msg"`
	Path      string                               `json:"path"`
	RequestID string                               `json:"requestId"`
	Data      BehavioralVerificationVerifyRespData `json:"data"`
}

type BehavioralVerificationVerifyRespData struct {
	Result bool `json:"result"`
}

type BootstrapReq struct {
	AccessKey    string `form:"accessKey"`
	AccessSecret string `form:"accessSecret"`
}

type BootstrapResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type CallbackAliyuncsOssFrontendUploadReq struct {
	Filename  string `form:"filename,optional"`
	Height    int32  `form:"height,optional"`
	Width     int32  `form:"width,optional"`
	MimeType  string `form:"mimeType,optional"`
	Size      int64  `form:"size,optional"`
	UUID      string `path:"uuid,optional"`
	AccessKey string `path:"accessKey,optional"`
}

type CallbackAliyuncsOssFrontendUploadResp struct {
	Code      int32                                     `json:"code"`
	Msg       string                                    `json:"msg"`
	RequestID string                                    `json:"requestId"`
	Path      string                                    `json:"path"`
	Data      CallbackAliyuncsOssFrontendUploadRespData `json:"data"`
}

type CallbackAliyuncsOssFrontendUploadRespData struct {
	FileName    string `json:"filename"`
	Size        int64  `json:"size"`
	MineType    string `json:"mimeType"`
	Height      int64  `json:"height"`
	Width       int64  `json:"width"`
	ExternalUrl string `json:"externalUrl"`
	Cover       string `json:"cover"`
}

type CcsCommonEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type CcsTestMsgReq struct {
	Key string `json:"key"`
	Msg string `json:"msg"`
}

type CodesReq struct {
	Lang string `form:"lang,optional" url:"lang"`
	Svc  string `form:"svc,optional" url:"svc"`
}

type CodesResp struct {
	Code      int32         `json:"code"`
	Msg       string        `json:"msg"`
	Path      string        `json:"path"`
	RequestID string        `json:"requestId"`
	Data      CodesRespData `json:"data"`
}

type CodesRespData struct {
	List     []ModelCode `json:"list"`
	Total    int64       `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32       `json:"pageSize"`
}

type CtasCommonEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type CtasTestMsgReq struct {
	Msg string `json:"msg"`
}

type DownloadPublicKeyReq struct {
}

type DownloadPublicKeyResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type HealthzResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type KeyExchangeReq struct {
	AccessKey string `json:"accessKey"`
	PublicKey string `json:"publicKey"`
}

type KeyExchangeResp struct {
	Code      int32               `json:"code"`
	Msg       string              `json:"msg"`
	Path      string              `json:"path"`
	RequestID string              `json:"requestId"`
	Data      KeyExchangeRespData `json:"data"`
}

type KeyExchangeRespData struct {
	PublicKey string `json:"publicKey"`
	ExpireAt  int64  `json:"expireAt"`
}

type KmsAkcCreateKeychainReq struct {
	CertType string `json:"certType"`
	Name     string `json:"name"`
}

type KmsAkcCreateKeychainResp struct {
	Code      int32                        `json:"code"`
	Msg       string                       `json:"msg"`
	Path      string                       `json:"path"`
	RequestID string                       `json:"requestId"`
	Data      KmsAkcCreateKeychainRespData `json:"data"`
}

type KmsAkcCreateKeychainRespData struct {
	Status     string `json:"status"`
	Name       string `json:"name"`
	SignMethod string `json:"signMethod"`
	CertType   string `json:"certType"`
	PublicKey  string `json:"publicKey"`
}

type KmsAkcGetKeychainPublicKeyReq struct {
	Name string `json:"name"`
}

type KmsAkcGetKeychainPublicKeyResp struct {
	Code      int32                              `json:"code"`
	Msg       string                             `json:"msg"`
	Path      string                             `json:"path"`
	RequestID string                             `json:"requestId"`
	Data      KmsAkcGetKeychainPublicKeyRespData `json:"data"`
}

type KmsAkcGetKeychainPublicKeyRespData struct {
	Name       string `json:"name"`
	SignMethod string `json:"signMethod"`
	CertType   string `json:"certType"`
	PublicKey  string `json:"publicKey"`
}

type KmsAkcSignReq struct {
	Name        string `json:"name"`
	SignContent string `json:"signContent"`
}

type KmsAkcSignResp struct {
	Code      int32              `json:"code"`
	Msg       string             `json:"msg"`
	Path      string             `json:"path"`
	RequestID string             `json:"requestId"`
	Data      KmsAkcSignRespData `json:"data"`
}

type KmsAkcSignRespData struct {
	Name string `json:"name"`
	Sign string `json:"sign"`
}

type KmsAkcVerifyReq struct {
	Name        string `json:"name"`
	SignContent string `json:"signContent"`
	Sign        string `json:"sign"`
}

type KmsAkcVerifyResp struct {
	Code      int32                `json:"code"`
	Msg       string               `json:"msg"`
	Path      string               `json:"path"`
	RequestID string               `json:"requestId"`
	Data      KmsAkcVerifyRespData `json:"data"`
}

type KmsAkcVerifyRespData struct {
	Name   string `json:"name"`
	Verify bool   `json:"verify"`
}

type KmsCommonEncryptResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type KmsSkcBatchDecryptReq struct {
	Data map[string]KmsSkcBatchDecryptReqItem `json:"data"`
}

type KmsSkcBatchDecryptReqItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcBatchDecryptResp struct {
	Code      int32                                 `json:"code"`
	Msg       string                                `json:"msg"`
	Path      string                                `json:"path"`
	RequestID string                                `json:"requestId"`
	Data      map[string]KmsSkcBatchDecryptRespItem `json:"data"`
}

type KmsSkcBatchDecryptRespItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
	Status   string `json:"status"`
}

type KmsSkcBatchEncryptReq struct {
	Data map[string]KmsSkcBatchEncryptReqItem `json:"data"`
}

type KmsSkcBatchEncryptReqItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcBatchEncryptResp struct {
	Code      int32                                 `json:"code"`
	Msg       string                                `json:"msg"`
	Path      string                                `json:"path"`
	RequestID string                                `json:"requestId"`
	Data      map[string]KmsSkcBatchEncryptRespItem `json:"data"`
}

type KmsSkcBatchEncryptRespItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
	Status   string `json:"status"`
}

type KmsSkcCompareItem struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcCompareReq struct {
	List []KmsSkcCompareItem `json:"list"`
}

type KmsSkcCompareResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcCompareRespData `json:"data"`
}

type KmsSkcCompareRespData struct {
	Status      string                      `json:"status"`
	CompareData bool                        `json:"compareData"`
	List        []KmsSkcCompareRespDataItem `json:"list"`
}

type KmsSkcCompareRespDataItem struct {
	Status   string `json:"status"`
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcCreateKeychainReq struct {
	Algorithm string `json:"algorithm"`
	Name      string `json:"name"`
}

type KmsSkcCreateKeychainResp struct {
	Code      int32                        `json:"code"`
	Msg       string                       `json:"msg"`
	Path      string                       `json:"path"`
	RequestID string                       `json:"requestId"`
	Data      KmsSkcCreateKeychainRespData `json:"data"`
}

type KmsSkcCreateKeychainRespData struct {
	Status string `json:"status"`
	Name   string `json:"name"`
}

type KmsSkcDecryptReq struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcDecryptResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcDecryptRespData `json:"data"`
}

type KmsSkcDecryptRespData struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcEncryptReq struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsSkcEncryptResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	Path      string                `json:"path"`
	RequestID string                `json:"requestId"`
	Data      KmsSkcEncryptRespData `json:"data"`
}

type KmsSkcEncryptRespData struct {
	Name     string `json:"name"`
	BaseData string `json:"baseData"`
}

type KmsTestMsgReq struct {
	Msg string `json:"msg"`
}

type ModelBucket struct {
	Id                     uint32 `json:"id"`
	CreatedAtUnix          int64  `json:"createdAtUnix"`
	UpdatedAtUnix          int64  `json:"updatedAtUnix"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	Region                 string `json:"region"`
	BucketInternetDomain   string `json:"bucketInternetDomain"`
	BucketInternalDomain   string `json:"bucketInternalDomain"`
	IsAccelerate           int32  `json:"isAccelerate"`
	BucketAccelerateDomain string `json:"bucketAccelerateDomain"`
	StaticDomain           string `json:"staticDomain"`
	CdnDomain              string `json:"cdnDomain"`
}

type ModelCode struct {
	Code string `json:"code"`
	Num  int32  `json:"num"`
	Msg  string `json:"msg"`
}

type ModelSasFile struct {
	Id              uint32   `json:"id"`
	CreatedAtUnix   int64    `json:"createdAtUnix"`
	UpdatedAtUnix   int64    `json:"updatedAtUnix"`
	CollectionId    uint32   `json:"collectionId"` // 文件集合ID
	Key             string   `json:"key"`          // 配置的key
	UploadType      int32    `json:"uploadType"`
	Bucket          string   `json:"bucket"`
	Endpoint        string   `json:"endpoint"`
	FileType        int32    `json:"fileType"`
	Uuid            string   `json:"uuid"`
	FileName        string   `json:"fileName"`
	FileExt         string   `json:"fileExt"`
	FileSize        int64    `json:"fileSize"`
	FormatName      string   `json:"formatName"`
	Remark          string   `json:"remark"`
	Path            string   `json:"path"`
	ExternalPath    string   `json:"externalPath"`
	MimeType        string   `json:"mimeType"`
	Width           int32    `json:"width"`
	Height          int32    `json:"height"`
	Cover           string   `json:"cover"`
	AlternativeCove []string `json:"alternativeCover"`
	UploadSignBody  string   `json:"uploadSignBody"`
}

type ModelSasFolder struct {
	Id            uint32 `json:"id"`
	CreatedAtUnix int64  `json:"createdAtUnix"`
	UpdatedAtUnix int64  `json:"updatedAtUnix"`
	Name          string `json:"name"`
	Alias         string `json:"alias"`
	Cover         string `json:"cover"`
	ArchiveType   int32  `json:"archiveType"`
	Owner         string `json:"owner"`
	Status        int32  `json:"status"`
	ParentId      uint32 `json:"parentId"`
}

type ModelZone struct {
	Label string `json:"label"`
	Code  string `json:"code"`
	Area  string `json:"area"`
}

type PingReq struct {
	Ping string `form:"ping"`
}

type PingResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type SasBootstrapReq struct {
}

type SasBootstrapResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
	Path      string `json:"path"`
	Data      string `json:"data"`
}

type SasFileApiCheckoutResultReq struct {
	FileId   uint32 `json:"fileId,optional"`
	FileUuid string `json:"fileUuid,optional"`
}

type SasFileApiCheckoutResultResp struct {
	Code      int32                            `json:"code"`
	Msg       string                           `json:"msg"`
	Path      string                           `json:"path"`
	RequestID string                           `json:"requestId"`
	Data      SasFileApiCheckoutResultRespData `json:"data"`
}

type SasFileApiCheckoutResultRespData struct {
	Status      string `json:"status"`
	Name        string `json:"name"`
	FolderId    uint32 `json:"folderId"`
	ExternalUrl string `json:"externalUrl"`
	Cover       string `json:"cover"`
	Remark      string `json:"remark"`
}

type SasFileApiCreateResp struct {
	Code      int32               `json:"code"`
	Msg       string              `json:"msg"`
	RequestID string              `json:"requestId"`
	Path      string              `json:"path"`
	Data      SasFileApiJsonIdReq `json:"data"`
}

type SasFileApiFormIdReq struct {
	Id uint32 `form:"id"`
}

type SasFileApiJsonIdReq struct {
	Id uint32 `json:"id"`
}

type SasFileApiJsonIdsReq struct {
	Ids []uint32 `json:"id"`
}

type SasFileApiOKResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
	Path      string `json:"path"`
	Data      string `json:"data"`
}

type SasFileCommonQueryListResp struct {
	Code      int32                          `json:"code"`
	Msg       string                         `json:"msg"`
	Path      string                         `json:"path"`
	RequestID string                         `json:"requestId"`
	Data      SasFileCommonQueryListRespData `json:"data"`
}

type SasFileCommonQueryListRespData struct {
	List     []ModelSasFile `json:"list"`
	Total    int64          `json:"total"`
	Page     int32          `json:"page"`
	PageSize int32          `json:"pageSize"`
}

type SasFileCommonQueryResp struct {
	Code      int32        `json:"code"`
	Msg       string       `json:"msg"`
	Path      string       `json:"path"`
	RequestID string       `json:"requestId"`
	Data      ModelSasFile `json:"data"`
}

type SasFileCommonSearchParams struct {
	Page           int32  `json:"page,optional"`
	PageSize       int32  `json:"pageSize,optional"`
	StartCreatedAt int64  `json:"startCreatedAt,optional"`
	EndCreatedAt   int64  `json:"endCreatedAt,optional"`
	Keyword        string `json:"keyword,optional"`
	Status         int32  `json:"status,optional"`
}

type SasFileCreateAndDirectUploadResp struct {
	Code      int32                                `json:"code"`
	Msg       string                               `json:"msg"`
	Path      string                               `json:"path"`
	RequestID string                               `json:"requestId"`
	Data      SasFileCreateAndDirectUploadRespData `json:"data"`
}

type SasFileCreateAndDirectUploadRespData struct {
	Id     uint32            `json:"id"`
	Url    string            `json:"url"`
	Fields map[string]string `json:"fields"`
}

type SasFolderApiCreateResp struct {
	Code      int32                 `json:"code"`
	Msg       string                `json:"msg"`
	RequestID string                `json:"requestId"`
	Path      string                `json:"path"`
	Data      SasFolderApiJsonIdReq `json:"data"`
}

type SasFolderApiFormIdReq struct {
	Id uint32 `form:"id"`
}

type SasFolderApiJsonIdReq struct {
	Id uint32 `json:"id"`
}

type SasFolderApiJsonIdsReq struct {
	Ids []uint32 `json:"id"`
}

type SasFolderApiOKResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"requestId"`
	Path      string `json:"path"`
	Data      string `json:"data"`
}

type SasFolderCommonQueryListResp struct {
	Code      int32                            `json:"code"`
	Msg       string                           `json:"msg"`
	Path      string                           `json:"path"`
	RequestID string                           `json:"requestId"`
	Data      SasFolderCommonQueryListRespData `json:"data"`
}

type SasFolderCommonQueryListRespData struct {
	List     []ModelSasFolder `json:"list"`
	Total    int64            `json:"total"`
	Page     int32            `json:"page"`
	PageSize int32            `json:"pageSize"`
}

type SasFolderCommonQueryResp struct {
	Code      int32          `json:"code"`
	Msg       string         `json:"msg"`
	Path      string         `json:"path"`
	RequestID string         `json:"requestId"`
	Data      ModelSasFolder `json:"data"`
}

type SasFolderCommonSearchParams struct {
	Page           int32  `json:"page,optional"`
	PageSize       int32  `json:"pageSize,optional"`
	StartCreatedAt int64  `json:"startCreatedAt,optional"`
	EndCreatedAt   int64  `json:"endCreatedAt,optional"`
	Keyword        string `json:"keyword,optional"`
	Status         int32  `json:"status,optional"`
}

type SasPresignerHeadObjectReq struct {
	BucketKey string `json:"bucketKey"`
	Path      string `json:"path"`
}

type SasPresignerHeadObjectResp struct {
	Code      int32                          `json:"code"`
	Msg       string                         `json:"msg"`
	RequestID string                         `json:"requestId"`
	Path      string                         `json:"path"`
	Data      SasPresignerHeadObjectRespData `json:"data"`
}

type SasPresignerHeadObjectRespData struct {
	Url    string `json:"url"`
	Method string `json:"method"`
}

type SasPresignerUploadReq struct {
	BucketKey string `json:"bucketKey"`
	Path      string `json:"path"`
}

type SasPresignerUploadResp struct {
	Code      int32                      `json:"code"`
	Msg       string                     `json:"msg"`
	RequestID string                     `json:"requestId"`
	Path      string                     `json:"path"`
	Data      SasPresignerUploadRespData `json:"data"`
}

type SasPresignerUploadRespData struct {
	Url    string            `json:"url"`
	Fields map[string]string `json:"fields"`
}

type SasQueryBucketReq struct {
	BucketKey string `json:"bucketKey"`
}

type SasQueryBucketResp struct {
	Code      int32       `json:"code"`
	Msg       string      `json:"msg"`
	RequestID string      `json:"requestId"`
	Path      string      `json:"path"`
	Data      ModelBucket `json:"data"`
}

type SignResultModel struct {
	AccessToken      string `json:"accessToken"`
	AccessExpiresAt  int64  `json:"accessExpiresAt"`
	RefreshToken     string `json:"refreshToken"`
	RefreshExpiresAt int64  `json:"refreshExpiresAt"`
}

type SmsInitReq struct {
	Key     string `json:"key"`
	Service string `json:"service"`
	Type    string `json:"type"`
	Zone    string `json:"zone"`
	Phone   string `json:"phone"`
}

type SmsInitResp struct {
	Code      int32           `json:"code"`
	Msg       string          `json:"msg"`
	Path      string          `json:"path"`
	RequestID string          `json:"requestId"`
	Data      SmsInitRespData `json:"data"`
}

type SmsInitRespData struct {
	Status string `json:"status"`
}

type SmsVerifyReq struct {
	Service    string `json:"service"`
	Type       string `json:"type"`
	Zone       string `json:"zone"`
	Phone      string `json:"phone"`
	VerifyCode string `json:"verifyCode"`
}

type SmsVerifyResp struct {
	Code      int32             `json:"code"`
	Msg       string            `json:"msg"`
	Path      string            `json:"path"`
	RequestID string            `json:"requestId"`
	Data      SmsVerifyRespData `json:"data"`
}

type SmsVerifyRespData struct {
	Result bool `json:"result"`
}

type TestReq struct {
	Msg string `json:"msg"`
}

type TestResp struct {
	Code      int32  `json:"code"`
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	RequestID string `json:"requestId"`
	Data      string `json:"data"`
}

type ZonesReq struct {
	Lang string `form:"lang,optional" url:"lang"`
}

type ZonesResp struct {
	Code      int32         `json:"code"`
	Msg       string        `json:"msg"`
	Path      string        `json:"path"`
	RequestID string        `json:"requestId"`
	Data      ZonesRespData `json:"data"`
}

type ZonesRespData struct {
	List     []ModelZone `json:"list"`
	Total    int64       `json:"total"`
	Page     int32       `json:"page"`
	PageSize int32       `json:"pageSize"`
}
