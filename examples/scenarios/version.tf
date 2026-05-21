terraform {
  required_providers {
    chronosphere = {
      # Version used by "make install" to simplify development.
      version = "0.0.1-dev"
      source  = "local/chronosphereio/chronosphere"
    }
  }
}
