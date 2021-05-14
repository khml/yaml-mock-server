# YAML-Mock-Server
HTTP MOCK Server Configured in YAML

## Config File

```yaml
cfg:
  debug: false # Run Debug Mode. logging request if true
  noCache: false # set Cache-Control: no-cache header if true
  port: 3000
  public: false # allow access to public-dir (execute dir)
  browser:
    open: true # open browser when run yaml-mock-server if true
    openPath: / # open path if open browser

routes:
  - path: / # routing path
    file: home/index.html # path to return file

  - path: /index.css
    file: home/index.css
```