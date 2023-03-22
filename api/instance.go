package api

type Instance struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	IP              string   `json:"ip"`
	Status          string   `json:"status"`
	SshKeyNames     []string `json:"ssh_key_names"`
	FileSystemNames []string `json:"file_system_names"`
	Region          struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"region"`
	InstanceType struct {
		Name             string      `json:"name"`
		Description      string      `json:"description"`
		PriceCentsHourly int         `json:"price_cents_per_hour"`
		Specs            interface{} `json:"specs"`
	} `json:"instance_type"`
	Hostname     string `json:"hostname"`
	JupyterToken string `json:"jupyter_token"`
	JupyterUrl   string `json:"jupyter_url"`
}

type ListInstanceResponse struct {
	Data []Instance `json:"data"`
}

type InstanceGetAPIResponse struct {
	Data Instance `json:"data"`
}

type InstanceCreateAPIRequest struct {
	RegionName       string   `json:"region_name"`
	InstanceTypeName string   `json:"instance_type_name"`
	SSHKeyNames      []string `json:"ssh_key_names"`
	FileSystemNames  []string `json:"file_system_names,omitempty"`
	Quantity         int      `json:"Quantity"`
	Name             *string  `json:"name,omitempty"`
}

type InstanceCreateAPIResponse struct {
	Data struct {
		InstanceIds []string `json:"instance_ids"`
	} `json:"data"`
}

type InstanceDeleteApiRequest struct {
	InstanceIds []string `json:"instance_ids"`
}

type InstanceDeleteApiResponse struct {
	Data struct {
		TerminatedInstances []Instance `json:"terminated_instances"`
	} `json:"data"`
}
