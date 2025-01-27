package htmlreport

import (
	"strconv"
	"strings"

	"github.com/klouddb/klouddbshield/model"
)

type SectionProgress struct {
	SectionName string
	Score       int
	MaxScore    int
	Percentage  float64
	Color       string
	AnchorID    string
}

type SectionSummary struct {
	Data    []SectionProgress
	Overall SectionProgress
}

type PostgresReport struct {
	PostgresResults []*model.Result
	Summary         *SectionSummary
	PostgresVersion string
}

func (h *HtmlReportHelper) RegisterPostgresReportData(listOfResults []*model.Result, scoreMap map[int]*model.Status, database string, printSummary bool) {

	if !printSummary {
		// Add the data to the template
		tabData := &PostgresReport{
			PostgresResults: listOfResults,
			PostgresVersion: database,
		}
		h.AddTab("Postgres Security Report", tabData)
		return
	}

	// Define sectiions statically for demonstration
	sections := []*model.Section{
		{Name: "Overall Score", Score: 0, MaxScore: 0, Color: "#373854"},
		{Name: "Section 1  - Installation and Patches", Score: 0, MaxScore: 0, Color: "#EA4335"},
		{Name: "Section 2  - Directory and File Permissions", Score: 0, MaxScore: 0, Color: "#FBBC05"},
		{Name: "Section 3  - Logging Monitoring and Auditing", Score: 0, MaxScore: 0, Color: "#34A853"},
		{Name: "Section 4  - User Access and Authorization", Score: 0, MaxScore: 0, Color: "#673AB7"},
		{Name: "Section 5  - Connection and Login", Score: 0, MaxScore: 0, Color: "#4285F4"},
		{Name: "Section 6  - Postgres Settings", Score: 0, MaxScore: 0, Color: "#9E379F"},
		{Name: "Section 7  - Replication", Score: 0, MaxScore: 0, Color: "#7BB3FF"},
		{Name: "Section 8  - Special Configuration Considerations", Score: 0, MaxScore: 0, Color: "#FF6F69"},
	}

	// Find the first control in each section and it's id
	sectionLeaderMap := make(map[int]string)

	for _, result := range listOfResults {
		sectionId, _ := strconv.Atoi(strings.Split(result.Control, ".")[0])
		if sectionLeaderMap[sectionId] == "" {
			sectionLeaderMap[sectionId] = result.Control + result.Title
		}
	}

	data := &SectionSummary{}
	// Render the progress bars for each section, apart from overall section
	for idx, section := range sections[1:] {
		section.Score = scoreMap[idx+1].Pass
		section.MaxScore = scoreMap[idx+1].Pass + scoreMap[idx+1].Fail

		if section.MaxScore == 0 {
			continue
		}

		progressPercentage := float64(section.Score) / float64(section.MaxScore) * 100
		data.Data = append(data.Data, SectionProgress{
			SectionName: section.Name,
			Score:       section.Score,
			MaxScore:    section.MaxScore,
			Percentage:  progressPercentage,
			Color:       section.Color,
			AnchorID:    sectionLeaderMap[idx+1],
		})
	}

	// Render Overall Progress bar
	sections[0].Score = scoreMap[0].Pass
	sections[0].MaxScore = scoreMap[0].Pass + scoreMap[0].Fail
	overallPercentage := float64(sections[0].Score) / float64(sections[0].MaxScore) * 100
	data.Overall = SectionProgress{
		SectionName: sections[0].Name,
		Score:       sections[0].Score,
		MaxScore:    sections[0].MaxScore,
		Percentage:  overallPercentage,
		Color:       sections[0].Color,
	}

	// Add the data to the template
	tabData := &PostgresReport{
		PostgresResults: listOfResults,
		Summary:         data,
		PostgresVersion: database,
	}
	h.AddTab("Postgres Security Report", tabData)
}

func (h *HtmlReportHelper) RegisterMysqlReportData(listOfResults []*model.Result, scoreMap map[int]*model.Status) {
	// Define sectiions statically for demonstration

	sections := []*model.Section{
		{Name: "Overall Score", Score: 0, MaxScore: 0, Color: "#373854"},
		{Name: "Section 1  - Operating system", Score: 0, MaxScore: 0, Color: "#373854"},
		{Name: "Section 2  - Installation and Planning", Score: 0, MaxScore: 0, Color: "#EA4335"},
		{Name: "Section 3  - File Permissions", Score: 0, MaxScore: 0, Color: "#FBBC05"},
		{Name: "Section 4  - General", Score: 0, MaxScore: 0, Color: "#34A853"},
		{Name: "Section 5  - MySQL Permissions", Score: 0, MaxScore: 0, Color: "#673AB7"},
		{Name: "Section 6  - Auditing and Logging", Score: 0, MaxScore: 0, Color: "#4285F4"},
		{Name: "Section 7  - Authentication", Score: 0, MaxScore: 0, Color: "#9E379F"},
		{Name: "Section 8  - Network", Score: 0, MaxScore: 0, Color: "#7BB3FF"},
		{Name: "Section 9  - Replication", Score: 0, MaxScore: 0, Color: "#FF6F69"},
		{Name: "Section 10 - MySQL InnoDB Cluster / Group Replication", Score: 0, MaxScore: 0, Color: "#673AB7"},
	}

	// Find the first control in each section and it's id
	sectionLeaderMap := make(map[int]string)

	for _, result := range listOfResults {
		sectionId, _ := strconv.Atoi(strings.Split(result.Control, ".")[0])
		if sectionLeaderMap[sectionId] == "" {
			sectionLeaderMap[sectionId] = result.Control + result.Title
		}
	}

	data := &SectionSummary{}

	// Render the progress bars for each section, apart from overall section
	for idx, section := range sections[1:] {
		section.Score = scoreMap[idx+1].Pass
		section.MaxScore = scoreMap[idx+1].Pass + scoreMap[idx+1].Fail

		if section.MaxScore == 0 {
			continue
		}

		progressPercentage := float64(section.Score) / float64(section.MaxScore) * 100
		data.Data = append(data.Data, SectionProgress{
			SectionName: section.Name,
			Score:       section.Score,
			MaxScore:    section.MaxScore,
			Percentage:  progressPercentage,
			Color:       section.Color,
			AnchorID:    sectionLeaderMap[idx+1],
		})

	}

	// Render Overall Progress bar
	sections[0].Score = scoreMap[0].Pass
	sections[0].MaxScore = scoreMap[0].Pass + scoreMap[0].Fail
	overallPercentage := float64(sections[0].Score) / float64(sections[0].MaxScore) * 100
	data.Overall = SectionProgress{
		SectionName: sections[0].Name,
		Score:       sections[0].Score,
		MaxScore:    sections[0].MaxScore,
		Percentage:  overallPercentage,
		Color:       sections[0].Color,
	}

	// Add the data to the template
	tabData := &PostgresReport{
		PostgresResults: listOfResults,
		Summary:         data,
	}

	h.AddTab("Mysql", tabData)
}

func (h *HtmlReportHelper) RegisterHBAReportData(listOfResults []*model.HBAScannerResult) {
	h.AddTab("HBA Scanner Report", listOfResults)
}

func (h *HtmlReportHelper) RegisterUserlistData(listOfResults []model.UserlistResult) {
	h.AddTab("Users Report", listOfResults)
}
