from diagrams import Cluster, Diagram, Node
from diagrams.aws.compute import ECS, Lambda
from diagrams.aws.network import Route53, CloudFront, APIGateway
from diagrams.aws.storage import S3
from diagrams.aws.integration import SQS
from diagrams.aws.management import Cloudwatch
from diagrams.aws.general import User
from diagrams.aws.database import RDS, ElastiCache
from IPython.display import Image

with Diagram("Case Tecnico Itaú - PH", show=False, direction="LR"):

    user = User("Usuário Web")

    with Cluster("Entrada Frontend"):
        route53 = Route53("Route 53")
        cloudfront = CloudFront("CloudFront")
        s3_front = S3("SPA no S3")
        user >> route53 >> cloudfront >> s3_front

    with Cluster("API Gateway + Segurança"):
        api_gateway = APIGateway("API Gateway (com API Key)")
        auth_lambda = Lambda("Lambda Auth")
        cloudfront >> api_gateway >> auth_lambda

    with Cluster("Core do serviço"):
        alb = ALB("ALB - Balanceador de Carga")
        fargate = Fargate("Cluster Fargate")
        scaling = AutoScaling("Auto Scaling")
        api_gateway >> alb >> fargate >> scaling

        with Cluster("ECS - Serviços de Negócio"):
            usuarios = ECS("Usuários")
            filmes = ECS("Filmes")
            sessoes = ECS("Sessões")
            produtos = ECS("Produtos")
            ingressos = ECS("Ingressos")
            pedidos = ECS("Pedidos")
            pagamento = ECS("Pagamento")
            fargate >> [usuarios, filmes, sessoes, produtos, ingressos, pedidos]

    with Cluster("Fila de Pagamento com Idempotência"):
        fila_pagamento = SQS("SQS Pagamento")
        dlq_pagamento = SQS("DLQ Pagamento")
        pagamento_lambda = Lambda("Lambda Pagamento")
        pedidos >> fila_pagamento >> pagamento_lambda >> pagamento
        fila_pagamento >> dlq_pagamento

    with Cluster("Notificações Assíncronas"):
        fila_notif = SQS("SQS Notificações")
        lambda_notif = Lambda("Lambda Notificação")
        [pedidos, pagamento] >> fila_notif >> lambda_notif

    with Cluster("Persistência"):
        postgres = RDS("RDS PostgreSQL")
        redis = ElastiCache("Redis TTL/Assentos")
        for svc in [usuarios, filmes, sessoes, produtos, ingressos, pedidos, pagamento]:
            svc >> postgres
        ingressos >> redis
        sessoes >> redis

    with Cluster("Observabilidade"):
        cloudwatch = Cloudwatch("CloudWatch")
        for svc in [usuarios, filmes, sessoes, produtos, ingressos, pedidos,
                    pagamento, pagamento_lambda,
                    auth_lambda, lambda_notif]:
            svc >> cloudwatch


Image("case_tecnico_itaú_-_ph.png")