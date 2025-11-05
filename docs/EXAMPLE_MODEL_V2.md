## Example using model v2

```
{
  "type": "germline",
  "status_code": "in_progress",
  "project_code": "CBTN",
  "diagnostic_lab_code": "DGL123",
  "panel_code": "LEUKEMIA",
  "primary_condition_mondo_code": "MONDO:0001234",
  "request_priority_code": "routine",
  "patients": [
    {
      "affected_status_code": "affected",
      "family_history": [...],
      "observations_categorical": [...],
      "observations_text": [...],
      "organization_patient_id": "PA001", <-- external patient identifier for the proband
      "organization_code": "CHOP", <-- organization code for the proband
      "relation_to_proband_code": "proband"
    },
    {
      "affected_status_code": "not_affected",
      "family_history": [...],
      "observations_categorical": [...],
      "observations_text": [...],
      "organization_code": "CHOP",
      "organization_patient_id": "PA002",
      "relation_to_proband_code": "mother"
    }
  ],

  "sequencing_experiments": [
    { // Sequencing experiment for the proband
      "aliquot": "ALQ001",
      "sample_organization_code": "CHOP", <-- organization code for the proband
      "submitter_sample_id": "SA001", <-- external sample identifier for the proband
      "capture_kit": "KAPA",
      "experiment_code": "...",
      "is_paired_end": true,
      "performer_lab_code": "LAB001",
      "read_length": 100,
      "request_priority_code": "routine",
      "run_alias": "run_123",
      "run_date": "2020-09-19T14:00:00Z",
      "run_name": "Run 123",
      "status_code": "completed",

    }
    {  // Sequencing experiment for the mother
      "aliquot": "ALQ002",
      "sample_organization_code": "CHOP",
      "submitter_sample_id": "SA002",
      "capture_kit": "KAPA",
      "experiment_code": "...",
      "is_paired_end": true,
      "performer_lab_code": "LAB001",
      "read_length": 100,
      "request_priority_code": "routine",
      "run_alias": "run_123",
      "run_date": "2020-09-19T14:00:00Z",
      "run_name": "Run 123",
      "status_code": "completed"

    }
  ],
  "tasks": [
    {
      "type": "alignment",
      "aliquot": "ALQ001",
      "output_documents": [
        {
          "data_category_code": "genomic",
          "data_type_code": "alignment",
          "format_code": "cram",
          "hash": "9e107d9d372bb6826bd81d3542a419d6",
          "name": "FILE001.cram",
          "size": 123456,
          "url": "s3://bucket/prefix/FILE001.cram"
        }
      ],
      "pipeline_code": "..."
    },
    {
      "type": "calling_variants",
      "input_documents": [
        {
          "url": "s3://bucket/prefix/FILE001.cram"
        }
      ],
      "output_documents": [
        {
          "data_category_code": "genomic",
          "data_type_code": "variant_calls",
          "format_code": "vcf",
          "hash": "e4d909c290d0fb1ca068ffaddf22cbd0",
          "name": "FILE002.vcf",
          "size": 654321,
          "url": "s3://bucket/prefix/FILE002.gvcf"
        }
      ],
      "pipeline_code": "..."
    },
    {
      "type": "exomiser",
      "input_documents": [
        {
          "url": "s3://bucket/prefix/FILE002.gvcf"
        }
      ],
      "output_documents": [
        {
          "data_category_code": "genomic",
          "data_type_code": "exomiser",
          "format_code": "tsv",
          "hash": "e4d909c290d0fb1ca068ffaddf22cbd0",
          "name": "FILE003.csv",
          "size": 654321,
          "url": "s3://bucket/prefix/FILE003.csv"
        }
      ],
      "pipeline_code": "..."

    },
    {
      "type": "joint_genotyping",
      "input_documents": [
        {
          "url": "s3://bucket/prefix/FILE002.gvcf"
        },
        {
          "url": "s3://bucket/prefix/FILE004.gvcf"
        }
      ],
      "output_documents": [
        {
          "data_category_code": "genomic",
          "data_type_code": "snv",
          "format_code": "vcf",
          "hash": "e4d909c290d0fb1ca068ffaddf22cbd0",
          "name": "FILE005.vcf",
          "size": 654321,
          "url": "s3://bucket/prefix/FILE005.vcf"
        }
      ],
      "pipeline_code": "..."

    }

  ]

}
```