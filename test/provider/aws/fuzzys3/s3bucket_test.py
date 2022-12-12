import boto3
from datetime import datetime
import string, random

def create_fake_bucket():
    # Create fake session to be used against motoserver
    session = boto3.Session(
        aws_access_key_id="test",
        aws_secret_access_key="test",
        profile_name="local",
        region_name="us-east-1",
    )
    s3 = session.client(
        "s3", endpoint_url="http://localhost:5000"
    )
    date = datetime.today().strftime('%Y%m%d%H%M%S')
    random_suffix = ''.join(random.choice(string.ascii_uppercase + string.digits) for _ in range(10))
    bucket_name = f"testbucket{date}{random_suffix}"
    s3.create_bucket(Bucket=bucket_name)

    # Verify that the bucket was created successfully
    response = s3.list_buckets()
    bucket_names = [bucket['Name'] for bucket in response['Buckets']]
    assert bucket_name in bucket_names
    print("S3 bucket created successfully")

create_fake_bucket()
