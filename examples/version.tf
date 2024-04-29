terraform {
  required_providers {
    chronosphere = {
      # Version used by "make install" to simplify development.
      version = "0.0.1-dev"
      source  = "local/chronosphereio/chronosphere"

      # To use the registry, update the version above and uncomment
      # the source line below:
      # source = "tf-registry.chronosphere.io/chronosphere/chronosphere"
    }
  }
}
