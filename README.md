# Relay-_caching_example
example of relay server with caching 

Relay server is deployed over istio which relays the request and responce back and forth to the client.
Relay server also does the caching the request and stores the request locally.
Relay server has envoy proxy enjected by the Istio.

#Steps to run the Application :

1. start the minikube 
2. start Istio 
3. Run command "kubectl apply -f relay.yaml "
4. Open another wsl cmd and Run command "minikube tunnel"
