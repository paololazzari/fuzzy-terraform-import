import boto3
from datetime import datetime
import string, random
import sys

### Availability zone is not yet implemented by moto

# def create_fake_resource():
#     # Create fake session to be used against motoserver
#     session = boto3.Session(
#         aws_access_key_id="test",
#         aws_secret_access_key="test",
#         profile_name="local",
#         region_name="us-east-1",
#     )
#     ec2 = session.client("ec2", endpoint_url="http://localhost:5000")
#     response = ec2.modify_availability_zone_group(
#         GroupName="availability-zone",
#         OptInStatus="opted-in"
#     )

#     # Verify that the launch template was created successfully
#     assert "State" in response

#     print(sys.argv[0] + " passed!")


# create_fake_resource()
