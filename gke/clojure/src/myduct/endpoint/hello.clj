(ns myduct.endpoint.hello
  (:require [compojure.core :refer :all]))

(defn hello-endpoint [config]
  (context "/hello" []
    (GET "/" []
      "Hello, World!")))
