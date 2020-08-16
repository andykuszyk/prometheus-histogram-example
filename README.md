# Prometheus Histogram Example
This repo showcases a simple example of the use of Prometheus histograms to instrument events in a process. For more details, checkout the accompanying blog post at [https://andykuszyk.github.io](https://andykuszyk.github.io/2020-07-24-prometheus-histograms.html).

## Usage
To use this example, run Prometheus and the example application with `docker-compose up`. Open Prometheus at [http://localhost:9090](http://localhost:9090) and search for the metrics prefixed with `histogram_metric_`.
