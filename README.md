# Serverless Cloud Infrastructure on AWS (Terraform, Go, CI/CD)

This repository contains the **backend** for a serverless resume website hosted on **AWS S3**, distributed globally via **AWS CloudFront**, with **Cloudflare** for DNS management.  

---

## 🚀 Highlights
- Deployed a static resume site to S3 with CloudFront and ACM for HTTPS - AWS CLI and Cloudflare DNS for custom domain setup.
- Implemented a serverless visitor counter with API Gateway, Lambda (Go), and DynamoDB (on-demand).
- Built CI/CD pipeline via GitHub Actions and OIDC-authenticated IAM roles for secure frontend deployments.
- Managed multi-account access via IAM Identity Center (SSO) using custom permission sets.
- Defined IaC in Terraform, covering S3 bucket policies, CORS, cache invalidation, DynamoDB setup, and Lambda deployment.

---

## 🗂️ Related Repositories
- [**Frontend**](https://github.com/LaurTudo/aws-cloud-resume-challenge-frontend) – Static site on S3 + CloudFront + Cloudflare DNS.
- **Infrastructure** (private for now) – Terraform for AWS resources and Cloudflare DNS.  