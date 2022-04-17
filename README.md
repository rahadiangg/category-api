# Category API

Repo adalah RESTFful API biasa yang digunakan untuk ujicoba proses CI/CD Cloud Build di GCP yang nantinya akan di deploy pada CloudRun dan Kubernetes. Program ini dibuat dari tutorial punya bang Eko tentang golang RESTful API.

## Persiapan Deploy Cloud Run

- Service account CloudBuild tambahkan role ``Cloud Run Developer`` dan ``Service Account User``

## Perispan Deploy Kubernetes

- Service account CloudBuild tambahkan role ``Secret Manager Secret Accessor``
- Buat SSH Key terlebih dahulu (private & public key)
- Masukan private key ke Secret Manager di GCP
- Masukan public key ke repo [Category API Deployment](https://github.com/rahadiangg/category-api-deployment)
- 