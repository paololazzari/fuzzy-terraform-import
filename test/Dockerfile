FROM ubuntu

WORKDIR /test

# Install dependencies
RUN apt-get update -y && \
    apt-get install python3 -y && \
    apt-get install python3.10-venv -y && \
    apt-get install python3-pip -y

# Install virtual environment
RUN python3 -m venv env
RUN . env/bin/activate

# Install python dependecies
RUN pip install boto3 awscli awscli-plugin-endpoint moto flask flask-cors

# Create config file
RUN mkdir -p /root/.aws
COPY awsconfig.txt /root/.aws/config
COPY awscreds.txt /root/.aws/credentials
COPY script.sh .
COPY test_wrapper.sh .

RUN chmod +x script.sh
RUN chmod +x test_wrapper.sh

RUN apt-get install golang -y

ENTRYPOINT ["/test/script.sh"]