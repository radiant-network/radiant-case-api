package types

import "time"

type JsonArray[T any] []T

type OperationResponse struct {
	Status      string `json:"status_code" binding:"required"`
	OperationId string `json:"operation_id"`
}

type OperationErrorResponse struct {
	Status string                    `json:"status_code" binding:"required"`
	Id     string                    `json:"id"`
	Errors JsonArray[OperationError] `json:"errors"`
}

type OperationError struct {
	Field string `json:"field"`
	Error string `json:"error" binding:"required"`
	Code  string `json:"code"`
}

type ListCases struct {
	Cases JsonArray[Case] `json:"cases" binding:"required"`
}

type Case struct {
	ProjectCode               string                              `json:"project_code" binding:"required" example:"CBTN"`
	Type                      string                              `json:"type" binding:"required" enums:"germline,somatic" example:"germline"`
	StatusCode                string                              `json:"status_code" enums:"draft,in_progress,revoke,completed,incomplete,submitted,unknown" example:"in_progress"`
	RequestPriorityCode       string                              `json:"request_priority_code" enums:"routine,urgent,asap,stat" example:"routine"`
	DiagnosticLabCode         string                              `json:"diagnostic_lab_code" example:"DGL123"`
	PrimaryConditionMondoCode string                              `json:"primary_condition_mondo_code" example:"MONDO:0001234"`
	PanelCode                 string                              `json:"panel_code" example:"BREAST_CANCER_PANEL"`
	Patients                  JsonArray[CasePatient]              `json:"patients"`
	SequencingExperiments     JsonArray[CaseSequencingExperiment] `json:"sequencing_experiments"`
}

type PartialCase struct {
	StatusCode                string                              `json:"status_code" enums:"draft,in_progress,revoke,completed,incomplete,submitted,unknown" example:"completed"`
	DiagnosticLabCode         string                              `json:"diagnostic_lab_code" example:"DGL123"`
	PrimaryConditionMondoCode string                              `json:"primary_condition_mondo_code"`
	PanelCode                 string                              `json:"panel_code" example:"BREAST_CANCER_PANEL"`
	Patients                  JsonArray[CasePatient]              `json:"patients"`
	SequencingExperiments     JsonArray[CaseSequencingExperiment] `json:"sequencing_experiments"`
}

type CasePatient struct {
	OrganizationPatientId   string                                `json:"organization_patient_id" binding:"required" example:"PA000123"`
	OrganizationCode        string                                `json:"organization_code" binding:"required" example:"CHOP"`
	RelationToProbandCode   string                                `json:"relation_to_proband_code" binding:"required" enums:"proband,mother,father,brother,sister,unknown" example:"proband"`
	AffectedStatusCode      string                                `json:"affected_status_code" binding:"required" enums:"affected,not_affected,unknown" example:"affected"`
	ObservationsCategorical JsonArray[CaseObservationCategorical] `json:"observations_categorical"`
	ObservationsText        JsonArray[CaseObservationText]        `json:"observations_text"`
	FamilyHistory           JsonArray[CaseFamilyHistory]          `json:"family_history"`
}

type CaseFamilyHistory struct {
	FamilyMemberCode string `json:"family_member_code" binding:"required" enums:"father,mother,sibling,grandparent,uncle_aunt,cousin,child" example:"grandparent"`
	Condition        string `json:"condition" binding:"required" example:"breast cancer"`
}

type CaseObservationCategorical struct {
	ObservationCode    string `json:"type" binding:"required" enums:"phenotype,condition,ancestry,consanguinity" example:"phenotype"`
	CodeSystem         string `json:"code_system" binding:"required" example:"hpo"`
	CodeValue          string `json:"code_value" binding:"required" example:"HPO:0004322"`
	OnsetCode          string `json:"onset_code" enums:"unknown,antenatal,congenital,neonatal,infantile,childhood,juvenile,young_adult,middle_age,senior" example:"childhood"`
	InterpretationCode string `json:"interpretation_code" enums:"positive,negative" example:"positive"`
	Note               string `json:"note" example:"note"`
}

type CaseObservationText struct {
	ObservationCode string `json:"type" binding:"required" enums:"note"`
	Value           string `json:"value" binding:"required"`
}

type CaseSequencingExperiment struct {
	SubmitterSampleId      string                   `json:"submitter_sample_id" binding:"required" example:"SA000123"`
	SampleOrganizationCode string                   `json:"sample_organization_code" binding:"required" example:"CHOP"`
	ExperimentCode         string                   `json:"experiment_code" binding:"required"`
	StatusCode             string                   `json:"status_code" enums:"draft,in_progress,revoke,completed,incomplete,submitted,unknown" example:"completed"`
	RequestPriorityCode    string                   `json:"request_priority_code" enums:"routine,urgent,asap,stat" example:"routine"`
	Aliquot                string                   `json:"aliquot" binding:"required" example:"ALQ000123"`
	PerformerLabCode       string                   `json:"performer_lab_code" example:"LAB000123"`
	RunName                string                   `json:"run_name" example:"Run 123"`
	RunAlias               string                   `json:"run_alias" example:"run_123"`
	RunDate                time.Time                `json:"run_date" format:"date-time" example:"2020-09-19T14:00:00Z"`
	CaptureKit             string                   `json:"capture_kit" example:"KAPA"`
	IsPairedEnd            *bool                    `json:"is_paired_end" example:"true"`
	ReadLength             int                      `json:"read_length" example:"100"`
	Task                   SequencingExperimentTask `json:"task" binding:"required"`
}

type SequencingExperimentTask struct {
	Type          string              `json:"type" binding:"required" enums:"nga,toga,tnga,nea,toea,tnea,npa,topa,tnpa,tofpa,tra"`
	PipelineCode  string              `json:"pipeline_code" binding:"required"`
	NormalAliquot string              `json:"normal_aliquot" example:"ALQ000124"`
	Documents     JsonArray[Document] `json:"documents"`
}

type Document struct {
	Name             string `json:"name" binding:"required" example:"FILE000.cram"`
	DataCategoryCode string `json:"data_category_code" binding:"required" enums:"clinical,genomic" example:"genomic"`
	DataTypeCode     string `json:"data_type_code" binding:"required" enums:"alignment,snv,ssnv,gcnv,scnv,gsv,ssv,somfu,ssup,igv,cnvvis,exp,covgene,qcrun,exomiser,jointsnv" example:"alignment"`
	FormatCode       string `json:"format_code" binding:"required" enums:"cram,crai,gvcf,vcf,tbi,tgz,json,html,tsv,bw,bed,png,csv,pdf,txt" example:"cram"`
	Size             int64  `json:"size" binding:"required" example:"123456"`
	Url              string `json:"url" binding:"required" example:"s3://bucket/prefix/FILE000.cram"`
	Hash             string `json:"hash" binding:"required" example:"9e107d9d372bb6826bd81d3542a419d6"`
}

type TumorNormalTask struct {
	Type                  string              `json:"type" binding:"required" enums:"tnga,tnea,tnpa"`
	PipelineCode          string              `json:"pipeline_code" binding:"required"`
	OrganizationPatientId string              `json:"organization_patient_id" binding:"required" example:"PA000123"`
	OrganizationCode      string              `json:"organization_code" binding:"required" example:"CHOP"`
	NormalAliquot         string              `json:"normal_aliquot" example:"ALQ000124"`
	TumorAliquot          string              `json:"tumor_aliquot" example:"ALQ000123"`
	Documents             JsonArray[Document] `json:"documents"`
}

type UpdateSuccessResponse struct {
	Id string `json:"id"`
}

type UpdateErrorResponse struct {
	Id    string `json:"id"`
	Error string `json:"error" binding:"required"`
}

type ListPatients struct {
	Patients JsonArray[Patient] `json:"patients" binding:"required"`
}
type Patient struct {
	OrganizationPatientId  string `json:"organization_patient_id" binding:"required" example:"PA000123"`
	OrganizationCode       string `json:"organization_code" binding:"required" example:"CHOP"`
	OrganizationIdTypeCode string `json:"organization_type_code" binding:"required" example:"mrn"`
	IsAlive                *bool  `json:"is_alive" example:"true"`
}

type ListSamples struct {
	Samples JsonArray[Sample] `json:"samples" binding:"required"`
}
type Sample struct {
	OrganizationPatientId   string `json:"organization_patient_id" binding:"required" example:"PA000123"`
	OrganizationCode        string `json:"organization_code" binding:"required" example:"CHOP"`
	SubmitterSampleId       string `json:"submitter_sample_id" binding:"required" example:"SA000123"`
	SampleType              string `json:"sample_type" binding:"required" enums:"dna,rna,blood,solid_tissue" example:"blood"`
	TissueSite              string `json:"tissue_site"  example:"plasma"`
	Histology               string `json:"histology"  example:"normal" enums:"normal,tumoral"`
	ParentSubmitterSampleId string `json:"parent_submitter_sample_id" example:"SA000122"`
}
