Day of L e a r n i n g session: 1

The goal of this project is to become further aquainted with the Go programming
lanugage, and familiar with OpenShift.The goal is to deploy a webserver that
generates a gif of a Lissajous curve. The curve's parameters can be edited
through HTTP. This Go program will then be dockerized, pushed to my Quay
account, and deployed on an Openshift cluster.

LServer is the standalone Go program that hosts the image on port 1337.
GET on <LServer_IP>:1337/Params returns a JSON of the parameters.
PUT on <LServer_IP>:1337 modifies the curve's parameters.

Use updateLServer.sh and the example.json JSON file to modify the curves
parameters. Refresh the page after running to get the new curve image.
