#!/bin/bash

ENDPOINT="http://localhost:8000"
REGION="us-west-2"
MAX_ATTEMPTS=30

echo "Esperando a que DynamoDB Local esté disponible..."

for i in $(seq 1 $MAX_ATTEMPTS); do
    if aws dynamodb list-tables --endpoint-url $ENDPOINT --region $REGION &> /dev/null; then
        echo "DynamoDB Local está listo!"
        break
    fi

    if [ $i -eq $MAX_ATTEMPTS ]; then
        echo "Error: DynamoDB Local no respondió después de $MAX_ATTEMPTS intentos"
        exit 1
    fi

    echo "Intento $i/$MAX_ATTEMPTS - Esperando..."
    sleep 2
done

# Función para crear tabla con manejo de errores
create_table() {
    local table_name=$1
    echo "Creando tabla: $table_name"

    if aws dynamodb create-table \
        --table-name $table_name \
        --attribute-definitions AttributeName=_id,AttributeType=S \
        --key-schema AttributeName=_id,KeyType=HASH \
        --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
        --endpoint-url $ENDPOINT \
        --region $REGION 2>&1 | grep -q "ResourceInUseException"; then
        echo "⚠️  Tabla $table_name ya existe, omitiendo..."
    else
        echo "✅ Tabla $table_name creada exitosamente!"
    fi
}

# Crear todas las tablas
create_table "Users"
create_table "Installments"
create_table "Loans"

echo ""
echo "=== Tablas disponibles ==="
aws dynamodb list-tables --endpoint-url $ENDPOINT --region $REGION

echo ""
echo "✅ Setup completado!"
