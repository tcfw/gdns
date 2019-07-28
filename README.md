# gdns
GDNS Updates records in GCP Cloud DNS similar to dynamic DNS

# Usage
```sh
Usage:
  gdns [flags]

Flags:
  -d, --duration duration   Check duration (default 30s)
  -h, --help                help for gdns
  -p, --project string      GCP Project
  -r, --record string       GCP Managed Zone
  -v, --verbose             verbose
  -m, --zone string         GCP Managed Zone
```

# Example
```sh
gdns -p my-example-project -m example-dns-zone -r example.com. 
```

# Authentication
Either from the local gcloud auth files or from a JSON key references in the `GOOGLE_APPLICATION_CREDENTIALS` environment variable.
See https://cloud.google.com/docs/authentication/production#finding_credentials_automatically for more info.
