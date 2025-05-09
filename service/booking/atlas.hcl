data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

data "composite_schema" "app" {
  schema "booking" {
    url = "file://db/schema.hcl"
  }
  schema "booking" {
    url = data.external_schema.gorm.url
  }
}

env "local" {
  src = data.composite_schema.app.url
  url = "postgres://sa:PanPostGres1245@localhost:5432/evently_event?sslmode=disable"
  dev = "docker://postgres/15/dev"
  migration {
    dir = "file://db/migration"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}