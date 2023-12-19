# DevOps Utility Tool for IP and DNS Mapping

[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

The DevOps online utility tool accelerates cloud configuration debugging by mapping IP Addresses and their corresponding DNS values, offering visualization capabilities to understand IP-DNS connections effectively.

<img width="1209" alt="image" src="https://github.com/dev-sareno/ginamus/assets/33644608/c2bbb9a9-b61a-4a8b-9419-f97aff37e16a">

## Overview

This open-source project aims to simplify the process of resolving and visualizing IP Addresses and their linked DNS values, aiding in cloud configuration debugging and offering an intuitive method to understand the relationships between IPs and domains.

## Key Features

- **IP and DNS Mapping:** Resolve and visualize IP Addresses and DNS values effortlessly.
- **Cloud Debugging:** Accelerate the process of identifying server locations, webserver details, etc.
- **Visual Representation:** Mermaid.js powered graphs for a clear visual representation of connections.
- **Geolocation Integration:** Enrich the visualization by mapping IP Addresses to physical locations, enhancing geographic data representation.
- **Real-time Monitoring:** Continuous tracking of changes in IP-DNS mappings with alerts or notifications for alterations.

## Technology Stack

### Backend

- **Language:** Golang
- **Deployment:** Kubernetes
- **Database:** AWS DynamoDB

### Frontend

- **Framework:** ReactJS

## Current Features

The tool currently offers the following functionality:

- **Domain/Sub-Domain Resolution:** Accepts user-provided lists of domains/sub-domains, processes them, and visualizes the data using Mermaid.js.

## Planned Features

The roadmap for future enhancements includes:

- **Reverse Lookup:** Processing and visualization of lists of IP Addresses.
- **Combined Analysis:** Accepting both IP Addresses and domains/sub-domains for comprehensive analysis.
- **Export Options:** Generating PDF or image exports of visualizations.
- **Interactive Graphs:** Adding click functionality in Mermaid.js graphs for detailed information on nodes (IP or domain).
- **Historical Data Analysis:** Storing and analyzing historical mapping data to identify trends and anomalies over time.

## Getting Started

To get started with the project:
```shell
$ git clone https://github.com/dev-sareno/ginamus.git
$ cd ginamus
$ docker compose up -d
```

## Usage

1. **Provide a list of domains/sub-domains** through the provided UI.
2. **Visualize the IP and DNS mappings** using the generated Mermaid.js graphs.

## Contributing

We welcome contributions from the community! To contribute to the project:

- Read our [Code of Conduct](CODE_OF_CONDUCT.md).
- Review our [Contributing Guidelines](CONTRIBUTING.md).
- Fork the repository, make your changes, and submit a pull request.
- Ensure your code follows the project's coding conventions.

## Code of Conduct

This project follows our [Code of Conduct](CODE_OF_CONDUCT.md) to ensure a respectful and inclusive environment for everyone.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

This project acknowledges the generous support of JetBrains through their Open Source Support (OSS) program, providing a free license for their powerful GoLand IDE.

[<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.png" alt="GoLand logo">](https://www.jetbrains.com/go/)

### JetBrains Open Source Support

This project utilizes GoLand, an intelligent IDE for Go development, made possible by JetBrains' commitment to supporting open-source initiatives. The GoLand logo used above is owned by JetBrains.

We extend our sincere gratitude to JetBrains for their continuous support of the open-source community.

[JetBrains Open Source Support Program](https://www.jetbrains.com/community/opensource/)

## Contact

For any queries or support, feel free to contact the project maintainer.
