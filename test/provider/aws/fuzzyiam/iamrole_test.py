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
    iam = session.client("iam", endpoint_url="http://localhost:5000")
    response = iam.create_role(RoleName="role", AssumeRolePolicyDocument="policy")

    # Verify that the role was created successfully
    role = response["Role"]
    assert "RoleId" in role

    print(sys.argv[0] + " passed!")


create_fake_resource()
