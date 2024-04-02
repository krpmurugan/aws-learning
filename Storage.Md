# Storage

## SAN - [Storage Area Network]

A SAN is a dedicated high-speed computer network that provides block-level access to data storage. You can connect different storage devices to multiple servers. Some common uses of SANs include sharing data and improving the utilisation of storage systems. You can also use a SAN as a centralised data management and backup system. It improves the performance of storage-intensive applications, leading to reduced downtime and maintenance costs.

### Can you name different types of switches used in a SAN?

There are different types of switches used in SANs, and each comes with a unique use. Some of the most popular switches are Fibre Channel switches, iSCSI switches and FCoE switches. All of these switches serve different purposes. A Fibre Channel switch transfers high-speed data, while iSCSI switches are useful for IP-based storage networking. An FCoE switch is useful for converged storage and network traffic.

### What is the difference b/w SAN and NAS ?The basic difference between SAN and NAS , 

SAN is Fabric based and NAS is Ethernet based.
SAN - Storage Area Network  
NAS - Network attached Storage

NAS is file-level storage accessed over a network, making it ideal for file sharing and easy management. SAN, on the other hand, offers block-level storage and is known for its high performance and low latency, making it suitable for critical applications and databases.

### What is a zone in the context of a Storage Area Network?
A zone is a logical grouping of devices within a Storage Area Network. Zones are used to simplify management of the network by allowing administrators to group devices together based on their function or purpose. For example, all of the devices that are used for a particular application could be grouped together in a zone.

### HOW IS FIBER CHANNEL DIFFERENT FROM ISCSI?

Fiber Channel generally provides high performance and high availability for business-critical applications, usually in the corporate data center. (High-end)
In contrast, iSCSI is generally used to provide SANs for smaller regional or departmental data centers. (Mid-range)

iSCSI leverages TCP/IP and Ethernet networks, making it cost-effective and easier to set up. Fibre Channel, on the other hand, is a dedicated, high-performance protocol designed for low-latency, high-throughput storage environments, often used in enterprise data centers

### What is Array

Array is a group of Independent physical disks to configure any Volumes or RAID volumes.

### What is RAID, and why is it important in storage systems

RAID allows us to combine multiple disks to enhance redundancy and performance. RAID 1 mirrors data across two drives for redundancy, while RAID 0 stripes data for improved performance. RAID 5 offers a balance of both, with redundancy and performance benefits

### What is the difference between synchronous and asynchronous replication?

Synchronous replication and asynchronous replication are two methods of copying data between storage systems. Synchronous replication ensures that data is written to both the source and destination simultaneously, providing data consistency but potentially impacting performance. Asynchronous replication allows data to be written to the destination at a later time, offering better performance at the cost of potential data lag.

