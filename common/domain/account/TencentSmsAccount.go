package account

type TencentSmsAccount struct {
	/**
	 * api相关
	 */
	Url    string `json:"url"`
	Region string `json:"region"`
	/**
	 * 账号相关
	 */
	SecretId    string `json:"secretId"`
	SecretKey   string `json:"secretKey"`
	SmsSdkAppId string `json:"smsSdkAppId"`
	TemplateId  string `json:"templateId"`
	SignName    string `json:"signName"`

	//标识渠道商Id
	SupplierId int `json:"supplierId"`
	//标识渠道商名字
	SupplierName string `json:"supplierName"`
}
