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
    response = ec2.create_key_pair(KeyName="key-pair")

    # Verify that the key pair was created successfully
    assert "KeyName" in response

    print(sys.argv[0] + " passed!")


create_fake_resource()
