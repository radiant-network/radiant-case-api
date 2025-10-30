# Radiant API – Case Creation Use Cases

This directory contains a set of **step-by-step guides** that demonstrate how to create and manage **cases** using the **Radiant API**.  
Each document focuses on a different scenario for creating and updating **germline** or **somatic** cases.  

These examples illustrate:
- How to create patients and samples.
- How to associate sequencing experiments.
- How to update a case as new information becomes available.

---

## 📁 Available Guides

### 1. [Germline Family Case](./GERMLINE_FAMILY.md)
Use this guide when:
- You are creating a **family-based germline case** (e.g., trio or duo).
- You want to associate multiple family members and their samples.

It shows how to:
- Create patients and samples for the proband and relatives.
- Link samples using parent–child relationships.
- Add new family members (e.g., the father) later using the `PATCH /cases/{id}` endpoint.

---

### 2. [Germline Solo Case – Step by Step](./GERMLINE_SOLO_STEP_BY_STEP.md)
Use this guide when:
- You have information for **only one patient (the proband)**.
- Sample and sequencing information will be added **later**.

It shows how to:
- Create a minimal germline case with just the patient.
- Update the case later to include sample and sequencing experiment data.

---

### 3. [Somatic Case](./SOMATIC.md)
Use this guide when:
- You want to create a **somatic case** (e.g., tumor–normal analysis).
- A related germline case already exists in Radiant.

It shows how to:
- Create a somatic case for a proband.
- Define tumor and normal samples.
- Add sequencing experiments for tumor-only and tumor–normal analyses.


---

### 4. [Germline Case with an existing sequencing experiment from another case](./GERMLINE_EXISTING_SOMATIC.md)
Use this guide when:
- You want to create a **germline family case**.
- Sequencing experiment for proband already exists in another case.

It shows how to reuse an existing sequencing experiment while creating a new germline case for the family.

---

### 4. [Somatic Case with an existing sequencing experiment from another case](./SOMATIC_EXISTING_GERMLINE.md)
Use this guide when:
- You want to create a **somatic case**.
- Sequencing experiment for normal tissue exists in another case.

It shows how to reuse an existing sequencing experiment while creating a new somatic case.


