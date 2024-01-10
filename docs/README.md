# Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add lambdaping https://udhos.github.io/lambdaping

Update files from repo:

    helm repo update

Search lambdaping:

    helm search repo lambdaping -l --version ">=0.0.0"
    NAME                    CHART VERSION	APP VERSION	DESCRIPTION
    lambdaping/lambdaping	0.1.0        	0.0.0      	Install lambdaping helm chart into kubernetes.

To install the charts:

    helm install my-lambdaping lambdaping/lambdaping
    #            ^             ^          ^
    #            |             |           \________ chart
    #            |             |
    #            |              \___________________ repo
    #            |
    #             \_________________________________ release (chart instance installed in cluster)

To uninstall the charts:

    helm uninstall my-lambdaping

# Source

<https://github.com/udhos/lambdaping>
