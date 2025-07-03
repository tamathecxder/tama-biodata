package tamabiodata

type Biodata struct {
	Name, LinkedInURL, GithubURL string
	Origin                       string
	Age                          int
}

func GetBiodata() Biodata {
	tama := Biodata{
		Name:        "Yudistira Eka Pratama",
		Origin:      "Cianjur",
		Age:         21,
		LinkedInURL: "https://www.linkedin.com/in/yudistira-eka-pratama",
		GithubURL:   "https://github.com/tamathecxder",
	}

	return tama
}
