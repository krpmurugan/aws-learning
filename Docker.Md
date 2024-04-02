# Docker

### Docker

Docker is an open-source containerization platform. It is used to automate the deployment of any application, using lightweight, portable containers

### Name and explain the various Docker components.

- Docker Client: Performs Docker build pull and run operations to open up communication with the Docker Host. The Docker command then employs Docker API to call any queries to run.
- Docker Host: Contains Docker daemon, containers, and associated images. The Docker daemon establishes a connection with the Registry. The stored images are the type of metadata dedicated to containerized applications.
- Registry: This is where Docker images are stored. There are two of them, a public registry and a private one. Docker Hub and Docker Cloud are two public registries available for use by anyone.

### Containers

Containers are packages of software that contain all of the necessary elements to run in any environment. In this way, containers virtualize the operating system and run anywhere, from a private data center to the public cloud or even on a developer's personal laptop.

### Docker Images:

Docker images are read-only templates that contain instructions for creating a container.

DockerFile: It is a text file that has all commands which need to be run for building a given image.

### Docker Swarm?

Docker Swarm is a clustering and orchestration tool for Docker containers, allowing you to create and manage a cluster of Docker nodes as a single virtual system. Docker Swarm enables you to deploy and scale containerized applications across multiple hosts, providing high availability and load balancing

### Virtual machine
A virtual machine (VM) is like a copy of an actual physical computer. A virtual server operates in a multi-tenant environment, meaning that multiple VMs run on the same physical hardware. Contrarily, Containers sit on top of a physical server and its host OS. Every container shares the host OS kernel and the binaries and libraries to run the required application. Once you set up a VM, you can run a container within the VM. Containers and VMs are not mutually exclusive and can co-exist alongside.

### Virtualization

Hardware virtualization is a technique that allows a physical machine to run multiple virtual machines (VMs) on top of a hypervisor. A hypervisor is a software layer that manages the access and allocation of the hardware resources, such as CPU, memory, disk, and network, among the VMs. Each VM has its own operating system (OS) and applications, and is isolated from the others.

### Containerization 

Containerization is a technique that allows a physical machine to run multiple containers on top of an OS kernel. A container is a lightweight and portable unit of software that contains the code, libraries, dependencies, and configuration files needed to run an application. Unlike VMs, containers do not have their own OS, but share the same kernel and resources of the host machine.

### Hypervisor

A hypervisor, also known as Virtual Machine Monitor (VMM), is a software program that executes on a computer to manage multiple VMs.
A hypervisor allows multiple operating systems to run on a single physical machine.
Hyper-V is a virtualization platform developed by Microsoft for running multiple virtual machines (VMs) on a single physical machine.

### What is VMware and how it is used

VMware is a virtualization and cloud computing company that develops virtualization software. They produce virtualization, networking, and security management tools. VMware software allows you to run a virtual computer on your physical computer. It increases the efficiency of your computer system.

### ESXi 
ESXi is a Type 1 hypervisor, meaning it runs directly on system hardware without the need for an OS. Type 1 hypervisors are also referred to as bare-metal hypervisors because they run directly on hardware.

### What’s the difference between virtualization and containerization?

Virtualization is an abstract version of a physical machine, while containerization is the abstract version of an application.

| Aspects | Virtualization  | Containerization | 
| :---: | :---: | :---: | 
|-| This helps developers to run and host multiple OS on the hardware of a single physical server. |This helps developers to deploy multiple applications using the same operating system on a single virtual machine or server. |
|Isolation | Each VM runs its own guest operating system.  | Containers share the host operating system kernel.|
|Resource Usage| Each VM requires its own set of resources. | Containers are lightweight and share host resources | 
|Basic| Simulate separate virtual machines for each application on a single physical machine. | Contains multiple applications on a single physical machine. | 
|Operating System| With virtualization, there is a separate operating system for each application. | With containerization, all the applications share the same operating system. | 
|Guest OS	| Virtualization supports various operating systems on the same physical machine or server. | Containerization restricts guest OS, as Containers are packaged to work with the same operating system as of server. | 
|Deployment| Hypervisor is required to implement VMs. |A container engine is used to deploy an individual container.
|File Size	|The virtual machine has a larger file size.|Containers are light-weight executables.

### Docker Daemon

The Docker daemon (dockerd) is a background service that runs on the host operating system and manages Docker containers. It is responsible for creating, running, and monitoring containers, as well as handling interactions between the Docker client and the Docker engine

### What do you know about the Docker system prune?

It’s a command used to remove all stopped containers, unused networks, build caches, and dangling images. Prune is one of the most useful commands in Docker. The syntax is:  $ docker system prune

### How do Docker daemon and the Docker client communicate with each other?

You use a combination of Rest API, socket.IO, and TCP to facilitate communication.

### Can you tell what is the functionality of a hypervisor?

A hypervisor is a software that makes virtualization happen because of which is sometimes referred to as the Virtual Machine Monitor. This divides the resources of the host system and allocates them to each guest environment installed.
This means that multiple OS can be installed on a single host system. Hypervisors are of 2 types:

1. Native Hypervisor: This type is also called a Bare-metal Hypervisor and runs directly on the underlying host system which also ensures direct access to the host hardware which is why it does not require base OS.
2. Hosted Hypervisor: This type makes use of the underlying host operating system which has the existing OS installed

### What can you tell about Docker Compose?
It is a YAML file consisting of all the details regarding various services, networks, and volumes that are needed for setting up the Docker-based application. So, docker-compose is used for creating multiple containers, host them and establish communication between them. For the purpose of communication amongst the containers, ports are exposed by each and every container.

### How will you ensure that a container 1 runs before container 2 while using docker compose?

Docker-compose does not wait for any container to be “ready” before going ahead with the next containers. In order to achieve the order of execution, we can use:

The “depends_on” which got added in version 2 of docker-compose can be used as shown in a sample docker-compose.yml file

