// O:\2016\P1\C2\2016042414_P1C2\L0\Ancillary\2016042414_P1C2_metadata.xml

type Flight struct {
	Description     string
	Collection_Date string
	Campaign        string
	Extraction
	Payload
	Sites_Flown
}

type Extraction struct {
	Software
	Extraction_date
	Extracted_By
	Notes
	Disk_Serial_Numbers
}

type Software struct {
	Rev           string
	URL           string
	Date          string
	LastChangedBy string
}

type Disk_Serial_Numbers struct {
	DCC                 string
	Camera              string
	GPSIMU              string
	Pony_Express_Backup string
}

type Payload struct {
	Spectrometer
	Lidar
	Camera
	GPSIMU
}

type Spectrometer struct {
	Serial_Number string
	Model         string
	Comments      string
	Manufacturer  string
}

type Lidar struct {
	Serial_Number string
	Model         string
	Comments      string
	Manufacturer  string
}

type Camera struct {
	Serial_Number string
	Model         string
	Comments      string
	Manufacturer  string
}

type GPSIMU struct {
	Serial_Number string
	Model         string
	Comments      string
	Manufacturer  string
}

type Sites_Flown struct {
	Site
}

type Site struct {
	SiteID     string
	DomainName string
	SiteType   string
	Northing   string
	Easting    string
	State      string
	Lat        string
	County     string
	Long       string
	PM_Code    string
	Science    string
	Domain     string
	Ownership  string
	Zone       string
}