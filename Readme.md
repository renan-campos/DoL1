<H3>Day of L e a r n i n g session: 1</h3>

The goal of this project is to become further aquainted with the Go programming
lanugage, and familiar with OpenShift.The goal is to deploy a webserver that
generates a gif of a Lissajous curve. The curve's parameters can be edited
through HTTP. This Go program will then be dockerized, pushed to my Quay
account, and deployed on an Openshift cluster.

LServer is the standalone Go program that hosts the image on port 1337.
- GET on LSERVER_IP:1337 returns the gif.<br/>
- GET on LSERVER_IP:1337/Params returns a JSON of the parameters.<br/>
- PUT on LSERVER_IP:1337/Params modifies the curve's parameters.<br/>

To modify the curve, edit the numbers in example.json and run:
$ curl -T example.json LSERVER_IP:1337
Refresh the page after running to get the new curve image.
