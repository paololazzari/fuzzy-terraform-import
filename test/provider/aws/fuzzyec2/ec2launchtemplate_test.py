import boto3
from datetime import datetime
import string, random
import sys


def create_fake_resource():
    # Create fake session to be used against motoserver
    session = boto3.Session(
        aws_access_key_id="test",
        aws_secret_access_key="test",
        profile_name="local",
        region_name="us-east-1",
    )
    ec2 = session.client("ec2", endpoint_url="http://localhost:5000")
    response = ec2.create_launch_template(
        LaunchTemplateName="launch-template",
        VersionDescription="launch template",
        LaunchTemplateData={"InstanceType": "t2.micro", "ImageId": "ami-123"},
    )

    # Verify that the launch template was created successfully
    response = ec2.describe_launch_templates(LaunchTemplateNames=["launch-template"])
    launch_template = response["LaunchTemplates"]
    assert len(launch_template) > 0

    print(sys.argv[0] + " passed!")


create_fake_resource()
