(ns myduct.endpoint.example
  (:require [compojure.core :refer :all]))

(defn example-endpoint [config]
  (context "/example" []
    (GET "/" []
      "This is an example endpoint")))
