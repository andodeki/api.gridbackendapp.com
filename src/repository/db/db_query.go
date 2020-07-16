package db

const (
	createUserQuery = `
	INSERT INTO users (
		email, password_hash, user_name, status
	)
	VALUES (
		:email, :password_hash, :user_name, :status
	)
	RETURNING user_id;
`
	getUserByIDQuery = `
	SELECT user_id, email,user_name,status, password_hash, created_at, deleted_at, updated_at
	FROM users 
	WHERE user_id = $1 AND deleted_at IS NULL;`

	getUserByEmailQuery = `
	SELECT user_id, email, user_name, status, password_hash, created_at, deleted_at, updated_at 
	FROM users 
	WHERE email = $1 AND deleted_at IS NULL;
`
)

const (
	createCarQuery = `
	INSERT INTO cars (
		user_id, car_value, car_make, car_model,year_of_manufacture,car_use,policy_period
	)
	VALUES (
		:user_id, :car_value, :car_make, :car_model,:year_of_manufacture,:car_use,:policy_period
	)
	RETURNING car_id;
	`
	updateCarQuery = `
	UPDATE cars
	SET car_value = :car_value,
		car_make = :car_make, 
		car_model = :car_model,
		year_of_manufacture = :year_of_manufacture,
		car_use = :car_use,
		policy_period = :policy_period
	WHERE car_id = :car_id;
	`
	getCarByIDQuery = `
	SELECT car_id, user_id, car_value, car_make, car_model, year_of_manufacture, car_use, policy_period
	FROM cars
	WHERE cars_id = $1;
	`
	listCarsByUserID = `
	SELECT car_id, user_id, car_value, car_make, car_model, year_of_manufacture, car_use, policy_period
	FROM cars
	WHERE user_id = $1 AND deleted_at IS NULL;
	`
	deleteCarQuery = `
	UPDATE cars
	SET deleted_at = NOW()
	WHERE car_id = $1 AND deleted_at IS NULL;
	`
)

const (
	createInsuranceDetailsQuery = `
	INSERT INTO insurance_details (
		user_id, car_id, insurance_type, basic_cover, add_ons,start_date
	)
	VALUES (
		:user_id, :car_id, :insurance_type, :basic_cover,:add_ons,:start_date
	)
	RETURNING insurance_id;
	`
	updateInsuranceDetailsQuery = `
	UPDATE insurance_details
	SET insurance_type = :insurance_type,
		basic_cover = :basic_cover, 
		add_ons = :add_ons,
		start_date = :start_date
	WHERE insurance_id = :insurance_id;
	`
	getInsuranceDetailsByIDQuery = `
	SELECT user_id, car_id, insurance_type, basic_cover, add_ons,start_date
	FROM insurance_details
	WHERE insurance_id = $1;
	`
	listInsuranceDetailsByUserID = `
	SELECT user_id, car_id, insurance_type, basic_cover, add_ons,start_date
	FROM insurance_details
	WHERE user_id = $1 AND deleted_at IS NULL;
	`
	listInsuranceDetailsByCarID = `
	SELECT user_id, car_id, insurance_type, basic_cover, add_ons,start_date
	FROM insurance_details
	WHERE car_id = $1 AND deleted_at IS NULL;
	`
	deleteInsuranceDetailsQuery = `
	UPDATE insurance_details
	SET deleted_at = NOW()
	WHERE insurance_id = $1 AND deleted_at IS NULL;
	`
)

const (
	createPaymentQuery = `
	INSERT INTO payments (
		user_id, car_id, insurance_id, payment_frequency, min_deposit, refund_amount, monthly_repayment,total_interest, payment_mode, currency, status
	)
	VALUES (
		:user_id, :car_id, :insurance_id, :payment_frequency, :min_deposit, :monthly_repayment,:total_interest, :payment_mode, :currency, :status
	)
	RETURNING payment_id;
	`
	updatePaymentQuery = `
	UPDATE payments
	SET payment_frequency = :payment_frequency,
		min_deposit = :min_deposit, 
		monthly_repayment = :monthly_repayment,
		total_interest = :total_interest,
		payment_mode = :payment_mode,
		currency = :currency,
		status = :status
	WHERE payments_id = :payments_id;
	`
	getPaymentByIDQuery = `
	SELECT user_id, car_id, insurance_id, payment_frequency, min_deposit, refund_amount, monthly_repayment,total_interest, payment_mode, currency, status
	FROM payments
	WHERE payment_id = $1;
	`
	listPaymentByUserID = `
	SELECT user_id, car_id, insurance_id, payment_frequency, min_deposit, refund_amount, monthly_repayment,total_interest, payment_mode, currency, status
	FROM payments
	WHERE user_id = $1 AND deleted_at IS NULL;
	`
	deletePaymentQuery = `
	UPDATE payment
	SET deleted_at = NOW(),
	refund_amount = :refund_amount
	WHERE payment_id = $1 AND deleted_at IS NULL;
	`
)
