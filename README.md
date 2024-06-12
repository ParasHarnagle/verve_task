# verve_task
### Basic Workflow of the application

![image](https://github.com/ParasHarnagle/verve_task/assets/117824030/9a58392e-35b4-49a2-bc17-e01ec9221a77)

This application comprises of gin based TLS server where is running concurrently in accordance with a cleanup task,
which is exectued after evry 30 minutes.
Redis storage is used to show case  an optmized approach for million records serving.
###### Assumption
In docker compose promotions.csv can be mapped from local volume to the service(verve) volume to be replaced after cleanup
A better approch can be placed with creating a fileserver to upload the promotions.csv
## Usage
This backend application can run via "docker compose".

```
docker compose up --build

```
Following curl commads can be called:
```
  curl https://localhost:1321/promotions/<promotion-id> --insecure
  
  curl https://localhost:1321/promotions/14603ea9-6d69-47cb-9a72-5d827cd5bdfc --insecure
```
##### insecure flag is passed so as to avoid generating the certs for the end user.


### Results
##### GET CALLS

![Screenshot 2024-06-12 at 8 17 33 PM](https://github.com/ParasHarnagle/verve_task/assets/117824030/f19d9588-b5be-4a93-b4c0-f798014600a7)
##### Server 
![Screenshot 2024-06-12 at 8 17 53 PM](https://github.com/ParasHarnagle/verve_task/assets/117824030/96898e01-5bd0-454d-b022-0ab5f243f174)

## How would you operate this app in production (e.g. deployment, scaling, monitoring)?
##### deployement
As i have implemented docker compose to integrate two services: server and caching
In basic production scenario assuming following services also to be present: frontend, database, fileserver, api-service, caching mechanism, auths etc.
So with respect to deployement I would use kuberenets services where all these services would be in an artifactory and appropriate k8s services, deployment,pv and other
k8s component would be integrated in an optimized desgin pattern methodolgy.
If native k8s i would be an extensive work i would select any major cloud provider such aws and select services as ECS, EKS for deployment.

Finally I would use helm chart approch to quickly deploy the application on the K8s platform
#### scaling
With respect to scaling, as I am using K8s i can leverage autoscaling or some scaling threshold can be controlled on the compute for Vertical scaling for hardware resources.
For horizontal scaling I can leverage "horizontal pod scaler" kind to implement on the deployments.

#### monitoring
K8s operator hub has services for monitoring such as prometheus,grafana and kibana which can be leveraged as per use even during production scenario.

## How would your application perform in peak periods (millions of requests per minute)? How would you optimize it?

Base of handling such requests at application would be caching and using of caching tools such as redis along wih adding a pooling mechanism to it.
I can also use horizontal auto scalers on deployments as per the demand and k8s based load balancer which can effectively distribute workloads on the basis of optmized algorithm such as a combination
of round robin and least connections type of algorithms.
Internal communicaiton within these microservice would be done via rpcs.
Finally a worker pool can implemented while server initialization which can drastically reduce the latency in serving.

## The .csv file could be very big (billions of entries) - how would your applicatioperform? How would you optimize it?
Application performance for billions of data would need multiple go routines searching the 'id' from the get call, thus would perform poorly.

For say billion records csv files, following strategy could be implemented:
A reader which reads and churns out 'n' of batches of file
This 'n' batches can be distributed to worker for concurrent processing were number of worker is dependent on cores of the system.
This chunk reading on the files or if any db presnet then on a sharded db would enhance the performance mutltiple times, dependent on the hardware resource.
If there is any methodolgy where this type of data can be map reduced then even better.



