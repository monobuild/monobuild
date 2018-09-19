package methods

const markerTemplate = `---

commands:
  - echo new marker
environment:
  MONOBUILD_VERSION: {{ .Version }}
label: {{ .Directory }}`
