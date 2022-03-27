# AWS Fargate Cloudformation 
This is an example of a Cloudformation template that deploys a container to AWS Fargate as a service.

Multiple AZs are used for high availability, SSL is terminated at the load balancer, health checks are used, and Auto Scaling keeps the CPU utilization at or below 50% (whatever % users set).

This includes but is not limited to:
* Cluster
* Cluster Security Group (associated with the Cluster)
* Load balancer (associated to HTTP Listener and Target Group)
* Load balancer Security Group (associated with the Load Balancer)
* HTTP Listener (associated with the load balancer)
* Service
* Target Group
* Task Definition
* Cloudwatch Log Group (associated with the Task Definition)
* ExecutionRole (used by Task Definition)
* TaskRole (used by Task Definition)
* Auto Scaling Role
* Auto Scaling Policy
* Auto Scaling Target



## Author
[Gift Banda](https://giftmbanda.com)