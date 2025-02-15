package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceGoogleKmsCryptoKey() *schema.Resource {
	dsSchema := datasourceSchemaFromResourceSchema(ResourceKMSCryptoKey().Schema)
	addRequiredFieldsToSchema(dsSchema, "name")
	addRequiredFieldsToSchema(dsSchema, "key_ring")

	return &schema.Resource{
		Read:   dataSourceGoogleKmsCryptoKeyRead,
		Schema: dsSchema,
	}

}

func dataSourceGoogleKmsCryptoKeyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	keyRingId, err := parseKmsKeyRingId(d.Get("key_ring").(string), config)
	if err != nil {
		return err
	}

	cryptoKeyId := KmsCryptoKeyId{
		KeyRingId: *keyRingId,
		Name:      d.Get("name").(string),
	}

	d.SetId(cryptoKeyId.CryptoKeyId())

	return resourceKMSCryptoKeyRead(d, meta)
}
