package iso8583

const (
	//Source: https://help.sap.com/doc/saphelp_pos22/2.2/en-US/e8/097fdfd4164e639361c7cbeaf306f7/content.htm?no_cache=true

	//ISO8583:1987 data elements
	BIT87_SECONDARY_BITMAP               = 1
	BIT87_PRIMARY_ACCOUNT_NUMBER         = 2
	BIT87_PROCESSING_CODE                = 3
	BIT87_TRANSACTION_AMOUNT             = 4
	BIT87_RECONCILIATION_AMOUNT          = 5
	BIT87_BILLING_AMOUNT                 = 6
	BIT87_TRANSMISSION_DATE_TIME         = 7
	BIT87_BILLING_FEE_AMOUNT             = 8
	BIT87_RECONCILIATION_CONVERSION_RATE = 9
	BIT87_BILLING_CONVERSION_RATE        = 10
	BIT87_SYSTEM_TRACE                   = 11
	BIT87_LOCAL_DATE_TIME                = 12
	BIT87_EFFECTIVE_DATE                 = 13
	BIT87_EXPIRY_DATE                    = 14
	BIT87_SETTLEMENT_DATE                = 15
	BIT87_CONVERSION_DATE                = 16
	BIT87_CAPTURE_DATE                   = 17
	BIT87_MERCHANT_TYPE                  = 18
	BIT87_ACQUIRER_COUNTRY_CODE          = 19
	BIT87_PAN_COUNTRY_CODE               = 20
	BIT87_FORWARD_COUNTRY_CODE           = 21
	BIT87_POS_DATA_CODE                  = 22
	BIT87_CARD_SEQ_NUMBER                = 23
	BIT87_FUNCTION_CODE                  = 24
	BIT87_MSG_REASON_CODE                = 25
	BIT87_ACCEPTOR_BUS_CODE              = 26
	BIT87_APPROVAL_CODE_LEN              = 27
	BIT87_RECONCILIATION_DATE            = 28
	BIT87_RECONCILIATION_INDICATOR       = 29
	BIT87_ORIGINAL_AMOUNT                = 30
	BIT87_ACQUIRER_REFERENCE_DATA        = 31
	BIT87_ACQUIRING_ID                   = 32
	BIT87_FORWARDING_ID                  = 33
	BIT87_EXTEND_PAN                     = 34
	BIT87_TRACK2                         = 35
	BIT87_TRACK3                         = 36
	BIT87_RETRIEVAL_NUMBER               = 37
	BIT87_APPROVAL_CODE                  = 38
	BIT87_ACTION_CODE                    = 39
	BIT87_SERVICE_CODE                   = 40
	BIT87_TERMINAL_ID                    = 41
	BIT87_MERCHANT_ID                    = 42
	BIT87_MERCHANT_NAME                  = 43
	BIT87_ADDITIONAL_RESPONSE_DATA       = 44
	BIT87_TRACK1                         = 45
	BIT87_AMOUNT_FEE                     = 46
	BIT87_ADDIDATA_NATIONAL              = 47
	BIT87_ADDIDATA_PRIVATE               = 48
	BIT87_TXN_CURRENCY_CODE              = 49
	BIT87_RECON_CURRENCY_CODE            = 50
	BIT87_BILLING_CURRENCY_CODE          = 51
	BIT87_ENCRYPTED_PIN                  = 52
	BIT87_SECURITY_CONTROL_INFO          = 53
	BIT87_ADDITIONAL_AMOUNT              = 54
	BIT87_IC_CARD_DATA                   = 55
	BIT87_ORIGINAL_DATA                  = 56
	BIT87_LIFE_CYCLE_CODE                = 57
	BIT87_AUTH_AGENT_CODE                = 58
	BIT87_TRANSPORT_DATA                 = 59
	BIT87_FIELD60                        = 60
	BIT87_FIELD61                        = 61
	BIT87_FIELD62                        = 62
	BIT87_FIELD63                        = 63
	BIT87_MAC1                           = 64
	BIT87_FIELD65                        = 65
	BIT87_ORIGINAL_FEE_AMOUNT            = 66
	BIT87_EXTEND_PAYMENT_DATE            = 67
	BIT87_RECEIVING_COUNTRY_CODE         = 68
	BIT87_SETTLEMENT_COUNTRY_CODE        = 69
	BIT87_AUTH_AGENT_COUNTRY_CODE        = 70
	BIT87_MESSAGE_NUMBER                 = 71
	BIT87_DATA_RECORD                    = 72
	BIT87_ACTION_DATE                    = 73
	BIT87_CREDIT_NUMBER                  = 74
	BIT87_CREDIT_REV_NUMBER              = 75
	BIT87_DEBIT_NUMBER                   = 76
	BIT87_DEBIT_REV_NUMBER               = 77
	BIT87_TRANSFER_NUMBER                = 78
	BIT87_TRANSFER_REV_NUMBER            = 79
	BIT87_INQUIEY_NUMBER                 = 80
	BIT87_AUTHORIZATION_NUMBER           = 81
	BIT87_INQUIRY_REV_NUMBER             = 82
	BIT87_PAYMENT_NUMBER                 = 83
	BIT87_PAYMENT_REV_NUMBER             = 84
	BIT87_FEE_COLLECTION_NUMBER          = 85
	BIT87_CREDIT_AMOUNT                  = 86
	BIT87_CREDIT_REV_AMOUNT              = 87
	BIT87_DEBIT_AMOUNT                   = 88
	BIT87_DEBIT_REV_AMOUNT               = 89
	BIT87_AUTHORIZATION_REV_NUMBER       = 90
	BIT87_TXN_DEST_COUNTRY_CODE          = 91
	BIT87_TXN_ORIG_COUNTRY_CODE          = 92
	BIT87_TXN_DEST_ID                    = 93
	BIT87_TXN_ORIG_ID                    = 94
	BIT87_ISSUER_REFER_DATA              = 95
	BIT87_KEY_MGMT_DATA                  = 96
	BIT87_NET_AMOUNT                     = 97
	BIT87_PAYEE                          = 98
	BIT87_SETTLEMENT_ID_CODE             = 99
	BIT87_RECEIVING_ID_CODE              = 100
	BIT87_FILE_NAME                      = 101
	BIT87_ACCNT_ID1                      = 102
	BIT87_ACCNT_ID2                      = 103
	BIT87_TXN_DESCRIPTION                = 104
	BIT87_CREDIT_CHARGEBACK_AMOUNT       = 105
	BIT87_DEBIT_CHARGEBACK_AMOUNT        = 106
	BIT87_CREDIT_CHARGEBACK_NUMBER       = 107
	BIT87_DEBIT_CHARGEBACK_NUMBER        = 108
	BIT87_CREDIT_FEE_AMOUNT              = 109
	BIT87_DEBIT_FEE_AMOUNT               = 110
	BIT87_FIELD111                       = 111
	BIT87_FIELD112                       = 112
	BIT87_FIELD113                       = 113
	BIT87_FIELD114                       = 114
	BIT87_FIELD115                       = 115
	BIT87_FIELD116                       = 116
	BIT87_FIELD117                       = 117
	BIT87_FIELD118                       = 118
	BIT87_FIELD119                       = 119
	BIT87_FIELD120                       = 120
	BIT87_FIELD121                       = 121
	BIT87_FIELD122                       = 122
	BIT87_FIELD123                       = 123
	BIT87_FIELD124                       = 124
	BIT87_FIELD125                       = 125
	BIT87_FIELD126                       = 126
	BIT87_FIELD127                       = 127
	BIT87_MAC2                           = 128

	//ISO8583:1993 data elements
	BIT93_SECONDARY_BITMAP               = 1
	BIT93_PRIMARY_ACCOUNT_NUMBER         = 2
	BIT93_PROCESSING_CODE                = 3
	BIT93_TRANSACTION_AMOUNT             = 4
	BIT93_RECONCILIATION_AMOUNT          = 5
	BIT93_BILLING_AMOUNT                 = 6
	BIT93_TRANSMISSION_DATE_TIME         = 7
	BIT93_BILLING_FEE_AMOUNT             = 8
	BIT93_RECONCILIATION_CONVERSION_RATE = 9
	BIT93_BILLING_CONVERSION_RATE        = 10
	BIT93_SYSTEM_TRACE                   = 11
	BIT93_LOCAL_DATE_TIME                = 12
	BIT93_EFFECTIVE_DATE                 = 13
	BIT93_EXPIRY_DATE                    = 14
	BIT93_SETTLEMENT_DATE                = 15
	BIT93_CONVERSION_DATE                = 16
	BIT93_CAPTURE_DATE                   = 17
	BIT93_MERCHANT_TYPE                  = 18
	BIT93_ACQUIRER_COUNTRY_CODE          = 19
	BIT93_PAN_COUNTRY_CODE               = 20
	BIT93_FORWARD_COUNTRY_CODE           = 21
	BIT93_POS_DATA_CODE                  = 22
	BIT93_CARD_SEQ_NUMBER                = 23
	BIT93_FUNCTION_CODE                  = 24
	BIT93_MSG_REASON_CODE                = 25
	BIT93_ACCEPTOR_BUS_CODE              = 26
	BIT93_APPROVAL_CODE_LEN              = 27
	BIT93_RECONCILIATION_DATE            = 28
	BIT93_RECONCILIATION_INDICATOR       = 29
	BIT93_ORIGINAL_AMOUNT                = 30
	BIT93_ACQUIRER_REFERENCE_DATA        = 31
	BIT93_ACQUIRING_ID                   = 32
	BIT93_FORWARDING_ID                  = 33
	BIT93_EXTEND_PAN                     = 34
	BIT93_TRACK2                         = 35
	BIT93_TRACK3                         = 36
	BIT93_RETRIEVAL_NUMBER               = 37
	BIT93_APPROVAL_CODE                  = 38
	BIT93_ACTION_CODE                    = 39
	BIT93_SERVICE_CODE                   = 40
	BIT93_TERMINAL_ID                    = 41
	BIT93_MERCHANT_ID                    = 42
	BIT93_MERCHANT_NAME                  = 43
	BIT93_ADDITIONAL_RESPONSE_DATA       = 44
	BIT93_TRACK1                         = 45
	BIT93_AMOUNT_FEE                     = 46
	BIT93_ADDIDATA_NATIONAL              = 47
	BIT93_ADDIDATA_PRIVATE               = 48
	BIT93_TXN_CURRENCY_CODE              = 49
	BIT93_RECON_CURRENCY_CODE            = 50
	BIT93_BILLING_CURRENCY_CODE          = 51
	BIT93_ENCRYPTED_PIN                  = 52
	BIT93_SECURITY_CONTROL_INFO          = 53
	BIT93_ADDITIONAL_AMOUNT              = 54
	BIT93_IC_CARD_DATA                   = 55
	BIT93_ORIGINAL_DATA                  = 56
	BIT93_LIFE_CYCLE_CODE                = 57
	BIT93_AUTH_AGENT_CODE                = 58
	BIT93_TRANSPORT_DATA                 = 59
	BIT93_FIELD60                        = 60
	BIT93_FIELD61                        = 61
	BIT93_FIELD62                        = 62
	BIT93_FIELD63                        = 63
	BIT93_MAC1                           = 64
	BIT93_FIELD65                        = 65
	BIT93_ORIGINAL_FEE_AMOUNT            = 66
	BIT93_EXTEND_PAYMENT_DATE            = 67
	BIT93_RECEIVING_COUNTRY_CODE         = 68
	BIT93_SETTLEMENT_COUNTRY_CODE        = 69
	BIT93_AUTH_AGENT_COUNTRY_CODE        = 70
	BIT93_MESSAGE_NUMBER                 = 71
	BIT93_DATA_RECORD                    = 72
	BIT93_ACTION_DATE                    = 73
	BIT93_CREDIT_NUMBER                  = 74
	BIT93_CREDIT_REV_NUMBER              = 75
	BIT93_DEBIT_NUMBER                   = 76
	BIT93_DEBIT_REV_NUMBER               = 77
	BIT93_TRANSFER_NUMBER                = 78
	BIT93_TRANSFER_REV_NUMBER            = 79
	BIT93_INQUIEY_NUMBER                 = 80
	BIT93_AUTHORIZATION_NUMBER           = 81
	BIT93_INQUIRY_REV_NUMBER             = 82
	BIT93_PAYMENT_NUMBER                 = 83
	BIT93_PAYMENT_REV_NUMBER             = 84
	BIT93_FEE_COLLECTION_NUMBER          = 85
	BIT93_CREDIT_AMOUNT                  = 86
	BIT93_CREDIT_REV_AMOUNT              = 87
	BIT93_DEBIT_AMOUNT                   = 88
	BIT93_DEBIT_REV_AMOUNT               = 89
	BIT93_AUTHORIZATION_REV_NUMBER       = 90
	BIT93_TXN_DEST_COUNTRY_CODE          = 91
	BIT93_TXN_ORIG_COUNTRY_CODE          = 92
	BIT93_TXN_DEST_ID                    = 93
	BIT93_TXN_ORIG_ID                    = 94
	BIT93_ISSUER_REFER_DATA              = 95
	BIT93_KEY_MGMT_DATA                  = 96
	BIT93_NET_AMOUNT                     = 97
	BIT93_PAYEE                          = 98
	BIT93_SETTLEMENT_ID_CODE             = 99
	BIT93_RECEIVING_ID_CODE              = 100
	BIT93_FILE_NAME                      = 101
	BIT93_ACCNT_ID1                      = 102
	BIT93_ACCNT_ID2                      = 103
	BIT93_TXN_DESCRIPTION                = 104
	BIT93_CREDIT_CHARGEBACK_AMOUNT       = 105
	BIT93_DEBIT_CHARGEBACK_AMOUNT        = 106
	BIT93_CREDIT_CHARGEBACK_NUMBER       = 107
	BIT93_DEBIT_CHARGEBACK_NUMBER        = 108
	BIT93_CREDIT_FEE_AMOUNT              = 109
	BIT93_DEBIT_FEE_AMOUNT               = 110
	BIT93_FIELD111                       = 111
	BIT93_FIELD112                       = 112
	BIT93_FIELD113                       = 113
	BIT93_FIELD114                       = 114
	BIT93_FIELD115                       = 115
	BIT93_FIELD116                       = 116
	BIT93_FIELD117                       = 117
	BIT93_FIELD118                       = 118
	BIT93_FIELD119                       = 119
	BIT93_FIELD120                       = 120
	BIT93_FIELD121                       = 121
	BIT93_FIELD122                       = 122
	BIT93_FIELD123                       = 123
	BIT93_FIELD124                       = 124
	BIT93_FIELD125                       = 125
	BIT93_FIELD126                       = 126
	BIT93_FIELD127                       = 127
	BIT93_MAC2                           = 128

	//ISO8583:2003 data elements
	BIT03_SECONDARY_BITMAP                 = 1
	BIT03_PRIMARY_ACCOUNT_NUMBER           = 2
	BIT03_PROCESSING_CODE                  = 3
	BIT03_TRANSACTION_AMOUNT               = 4
	BIT03_RECONCILIATION_AMOUNT            = 5
	BIT03_BILLING_AMOUNT                   = 6
	BIT03_TRANSMISSION_DATE_TIME           = 7
	BIT03_BILLING_FEE_AMOUNT               = 8
	BIT03_RECONCILIATION_CONVERSION_RATE   = 9
	BIT03_BILLING_CONVERSION_RATE          = 10
	BIT03_SYSTEM_TRACE                     = 11
	BIT03_LOCAL_DATE_TIME                  = 12
	BIT03_EFFECTIVE_DATE                   = 13
	BIT03_EXPIRY_DATE                      = 14
	BIT03_SETTLEMENT_DATE                  = 15
	BIT03_CONVERSION_DATE                  = 16
	BIT03_CAPTURE_DATE                     = 17
	BIT03_ERROR_INDICATOR                  = 18
	BIT03_ACQUIRER_COUNTRY_CODE            = 19
	BIT03_PAN_COUNTRY_CODE                 = 20
	BIT03_TXN_LIFECYCLE_DATA               = 21
	BIT03_POS_DATA_CODE                    = 22
	BIT03_CARD_SEQ_NUMBER                  = 23
	BIT03_FUNCTION_CODE                    = 24
	BIT03_MSG_REASON_CODE                  = 25
	BIT03_MERCHANT_CATEGORY_CODE           = 26
	BIT03_POS_CAPABILITY                   = 27
	BIT03_RECONCILIATION_DATE              = 28
	BIT03_RECONCILIATION_INDICATOR         = 29
	BIT03_ORIGINAL_AMOUNT                  = 30
	BIT03_ACQUIRER_REFERENCE_DATA          = 31
	BIT03_ACQUIRING_ID                     = 32
	BIT03_FORWARDING_ID                    = 33
	BIT03_ECOMMERCE_DATA                   = 34
	BIT03_TRACK2                           = 35
	BIT03_TRACK3                           = 36
	BIT03_RETRIEVAL_NUMBER                 = 37
	BIT03_APPROVAL_CODE                    = 38
	BIT03_ACTION_CODE                      = 39
	BIT03_SERVICE_CODE                     = 40
	BIT03_TERMINAL_ID                      = 41
	BIT03_MERCHANT_ID                      = 42
	BIT03_MERCHANT_DATA                    = 43
	BIT03_ADDITIONAL_RESPONSE_DATA         = 44
	BIT03_TRACK1                           = 45
	BIT03_AMOUNT_FEE                       = 46
	BIT03_ADDIDATA_NATIONAL                = 47
	BIT03_ADDIDATA_PRIVATE                 = 48
	BIT03_VERIF_DATA                       = 49
	BIT03_FIELD50                          = 50
	BIT03_FIELD51                          = 51
	BIT03_ENCRYPTED_PIN                    = 52
	BIT03_SECURITY_CONTROL_INFO            = 53
	BIT03_ADDITIONAL_AMOUNT                = 54
	BIT03_IC_CARD_DATA                     = 55
	BIT03_ORIGINAL_DATA                    = 56
	BIT03_LIFE_CYCLE_CODE                  = 57
	BIT03_AUTH_AGENT_CODE                  = 58
	BIT03_TRANSPORT_DATA                   = 59
	BIT03_FIELD60                          = 60
	BIT03_FIELD61                          = 61
	BIT03_FIELD62                          = 62
	BIT03_FIELD63                          = 63
	BIT03_MAC1                             = 64
	BIT03_FIELD65                          = 65
	BIT03_ORIGINAL_FEE_AMOUNT              = 66
	BIT03_EXTEND_PAYMENT_DATE              = 67
	BIT03_FILE_TRANSFER_CONTROL            = 68
	BIT03_FILE_TRANSFER_CONTROL_DATA       = 69
	BIT03_FILE_TRANSFER_DESCRIPTION        = 70
	BIT03_FIELD71                          = 71
	BIT03_DATA_RECORD                      = 72
	BIT03_ACTION_DATE                      = 73
	BIT03_RECONCILIATION_DATA_PRIMARY      = 74
	BIT03_RECONCILIATION_DATA_SECONDARY    = 75
	BIT03_FIELD76                          = 76
	BIT03_FIELD77                          = 77
	BIT03_FIELD78                          = 78
	BIT03_FIELD79                          = 79
	BIT03_FIELD80                          = 80
	BIT03_FIELD81                          = 81
	BIT03_FIELD82                          = 82
	BIT03_FIELD83                          = 83
	BIT03_FIELD84                          = 84
	BIT03_FIELD85                          = 85
	BIT03_FIELD86                          = 86
	BIT03_FIELD87                          = 87
	BIT03_FIELD88                          = 88
	BIT03_FIELD89                          = 89
	BIT03_FIELD90                          = 90
	BIT03_FIELD91                          = 91
	BIT03_FIELD92                          = 92
	BIT03_TXN_DEST_ID                      = 93
	BIT03_TXN_ORIG_ID                      = 94
	BIT03_ISSUER_REFER_DATA                = 95
	BIT03_KEY_MGMT_DATA                    = 96
	BIT03_NET_RECONCILIATION_AMOUNT        = 97
	BIT03_PAYEE                            = 98
	BIT03_SETTLEMENT_ID_CODE               = 99
	BIT03_RECEIVING_ID_CODE                = 100
	BIT03_FILE_NAME                        = 101
	BIT03_ACCNT_ID1                        = 102
	BIT03_ACCNT_ID2                        = 103
	BIT03_TXN_DESCRIPTION                  = 104
	BIT03_FIELD105                         = 105
	BIT03_FIELD106                         = 106
	BIT03_FIELD107                         = 107
	BIT03_FIELD108                         = 108
	BIT03_RECONCILIATION_FEE_AMOUNT_CREDIT = 109
	BIT03_RECONCILIATION_FEE_AMOUNT_DEBIT  = 110
	BIT03_FIELD111                         = 111
	BIT03_FIELD112                         = 112
	BIT03_FIELD113                         = 113
	BIT03_FIELD114                         = 114
	BIT03_FIELD115                         = 115
	BIT03_FIELD116                         = 116
	BIT03_FIELD117                         = 117
	BIT03_FIELD118                         = 118
	BIT03_FIELD119                         = 119
	BIT03_FIELD120                         = 120
	BIT03_FIELD121                         = 121
	BIT03_FIELD122                         = 122
	BIT03_FIELD123                         = 123
	BIT03_FIELD124                         = 124
	BIT03_FIELD125                         = 125
	BIT03_FIELD126                         = 126
	BIT03_FIELD127                         = 127
	BIT03_MAC2                             = 128
)
