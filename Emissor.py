#!/usr/bin/env python

# Exemplo de emissor RabbitMQ em Python
import pika

# Usuario e Senha de administrador do RabbitMQ
credentials = pika.PlainCredentials('usuario', 'senha')

# Dados do servidor do RabbitMQ
parameters = pika.ConnectionParameters('hostname',
                                       5672,
                                       '/',
                                       credentials)

# Estabelecendo conexao com servidor RabbitMQ
connection = pika.BlockingConnection(parameters) 
    

# Obtendo canal de comunicacao    
channel = connection.channel()

# Declarando uma nova fila chamada "sas" 
channel.queue_declare(queue='sas')

# Publicando uma mensagem na fila "sas"
channel.basic_publish(exchange='', routing_key='sas', body='Estou no SAS!')

print("Mensagem enviada!")

# Encerrando conexao com o servidor
connection.close()