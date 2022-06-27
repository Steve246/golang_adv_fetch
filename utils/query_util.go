package utils

const (
	INSERT_CUSTOMER = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values (:id,:name,:address,:phone,:email,:saldo)`

	UPDATE_CUSTOMER = `update customer set name=:name, address=:address, phone=:phone, email=:email where id=:id`

	DELETE_CUSTOMER = `update customer set is_status=0 where id=$1`

	INSERT_CUSTOMER_PS = `insert into customer 
						(id,name,address,phone,email,saldo) 
						values ($1,$2,$3,$4,$5,$6)`

	UPDATE_CUSTOMER_PS = `update customer set name=$1, address=$2, phone=$3, email=$4 where id=$5`

	DELETE_CUSTOMER_PS_HD = `delete from customer where id=$1`
	DELETE_CUSTOMER_PS_SD = `update customer set is_status=0 where id=$1`

	//Ini bagian shop

	INSERT_SHOP = `insert into shop 
						(id,no_siup,name,address,phone) 
						values (:id,:no_siup,:name,:address,:phone)`

	INSERT_SHOP_PS = `insert into customer (id,no_siup,name,address,phone) values ($1,$2,$3,$4)`

	UPDATE_SHOP = `update shop set name=:name, no_siup=:no_siup, address=:address, phone=:phone where id=:id`

	DELETE_SHOP_PS_HD = `delete from shop where id=$1`
	// DELETE_SHOP_PS_SD = `update shop set is_status=0 where id=$1`
)
