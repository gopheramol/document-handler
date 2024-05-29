package model

type ScanResult struct {
	ScanID           string           `json:"scan_id"`
	ScanTime         string           `json:"scan_time"`
	FileInfo         FileInfo         `json:"file_info"`
	ScanStatus       string           `json:"scan_status"`
	DetectionSummary DetectionSummary `json:"detection_summary"`
	ScanDetails      ScanDetails      `json:"scan_details"`
}

type FileInfo struct {
	FileName string `json:"file_name"`
	FileSize int    `json:"file_size"`
	FileType string `json:"file_type"`
	MD5      string `json:"md5"`
	SHA256   string `json:"sha256"`
}

type DetectionSummary struct {
	TotalThreatsDetected int      `json:"total_threats_detected"`
	Threats              []Threat `json:"threats"`
}

type Threat struct {
	ThreatID          string `json:"threat_id"`
	ThreatName        string `json:"threat_name"`
	Severity          string `json:"severity"`
	DetectionMethod   string `json:"detection_method"`
	FilePath          string `json:"file_path"`
	Status            string `json:"status"`
	RemediationAction string `json:"remediation_action"`
}

type ScanDetails struct {
	EngineVersion       string   `json:"engine_version"`
	DefinitionsVersion  string   `json:"definitions_version"`
	ScanDurationSeconds int      `json:"scan_duration_seconds"`
	ScannedFiles        int      `json:"scanned_files"`
	SkippedFiles        int      `json:"skipped_files"`
	Errors              []string `json:"errors"`
}
