# prometheus-http-sd

Prometheus HTTP Service Discovery Webservice

Setup example to apply in prometheus configuration:

```
  - job_name: http-sd-scrap
    scrape_interval: 15s
    http_sd_configs:
    - url: http://172.22.0.1:9990/targets
```

After that, setup the config.json:

```
[
    {
        "targets": ["hostname.com:80"],
        "labels": {
            "__metrics_path__": "/metrics"
        }
    },

    {
        "targets": ["hostname.com:9090"],
        "labels": {
            "__metrics_path__": "/metrics"
        }
    },

    {
        "targets": ["hostname.com:9093"],
        "labels": {
            "__metrics_path__": "/metrics"
        }
    }
]
```
# prometheus-sd-http
