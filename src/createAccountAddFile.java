import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;
import java.security.spec.InvalidKeySpecException;
import java.util.Properties;
import com.hedera.sdk.account.HederaAccount;
import com.hedera.sdk.account.HederaAccountCreateDefaults;
import com.hedera.sdk.common.HederaAccountID;
import com.hedera.sdk.common.HederaDuration;
import com.hedera.sdk.common.HederaPrecheckResult;
import com.hedera.sdk.common.HederaKey.KeyType;
import com.hedera.sdk.common.HederaTransactionAndQueryDefaults;
import com.hedera.sdk.common.HederaTransactionReceipt;
import com.hedera.sdk.common.HederaTransactionStatus;
import com.hedera.sdk.common.Utilities;
import com.hedera.sdk.cryptography.HederaCryptoKeyPair;
import com.hedera.sdk.file.HederaFile;
import com.hedera.sdk.node.HederaNode;
import com.hedera.sdk.transaction.HederaTransactionResult;

import file.FileAppend;
import file.FileCreate;
import file.FileDelete;
import file.FileGetContents;
import file.FileGetInfo;
import file.FileUpdate;

public final class createAccountAddFile {
	public static void main(String[] args) throws Exception {
		// node details
		String nodeAddress = "testnet43.hedera.com";
		int nodePort = 80;
		// these are loaded from config.properties below
		long nodeAccountShard = 0;
		long nodeAccountRealm = 0;
		long nodeAccountNum = 0;
		// your account details
		long payAccountShard = 0;
		long payAccountRealm = 0;
		long payAccountNum = 0;
		// you public and private keys
		String pubKey = "";
		String privKey = "";
		// load application properties
		Properties applicationProperties = new Properties();
		InputStream propertiesInputStream = null;
		propertiesInputStream = new FileInputStream("config.properties");
		// load a properties file
		applicationProperties.load(propertiesInputStream);
		// get the node's account values
		nodeAccountShard = Long.parseLong(applicationProperties.getProperty("nodeAccountShard"));
		nodeAccountRealm = Long.parseLong(applicationProperties.getProperty("nodeAccountRealm"));
		nodeAccountNum = Long.parseLong(applicationProperties.getProperty("nodeAccountNum"));
		// get my public/private keys
		pubKey = applicationProperties.getProperty("pubkey");
		privKey = applicationProperties.getProperty("privkey");
		// get my account details
		payAccountShard = Long.parseLong(applicationProperties.getProperty("payAccountShard"));
		payAccountRealm = Long.parseLong(applicationProperties.getProperty("payAccountRealm"));
		payAccountNum = Long.parseLong(applicationProperties.getProperty("payAccountNum"));
		// setup defaults for transactions and Queries */
		HederaTransactionAndQueryDefaults txQueryDefaults = new HederaTransactionAndQueryDefaults();
		// default memo to attach to transactions
		txQueryDefaults.memo = "Hello Future";
		// setup the node we're communicating with from the properties loaded above
		txQueryDefaults.node = new HederaNode(nodeAddress, nodePort,
				new HederaAccountID(nodeAccountShard, nodeAccountRealm, nodeAccountNum));
		// setup the paying account ID (got from the properties loaded above)
		txQueryDefaults.payingAccountID = new HederaAccountID(payAccountShard, payAccountRealm, payAccountNum);
		// setup the paying key pair (got from properties loaded above)
		txQueryDefaults.payingKeyPair = new HederaCryptoKeyPair(KeyType.ED25519, pubKey, privKey);
		// txQueryDefaults.transactionValidDuration = new HederaDuration(120, 0);
		txQueryDefaults.transactionValidDuration.seconds = (long) 120;
		txQueryDefaults.transactionValidDuration.nanos = (int) 0;
		// instantiate a new account object
		HederaAccount myNewAccount = new HederaAccount();
		// set its default Transaction and Query parameters
		myNewAccount.txQueryDefaults = txQueryDefaults;
		// create a new key for my new account
		HederaCryptoKeyPair newAccountKey = new HederaCryptoKeyPair(KeyType.ED25519);
		// now, setup default for account creation
		HederaAccountCreateDefaults defaults = new HederaAccountCreateDefaults();
		// auto renew period in seconds and nanos
		defaults.autoRenewPeriodSeconds = 86400;
		defaults.autoRenewPeriodNanos = 0;
		// defaults.maxReceiveProxyFraction = 0;
		// defaults.proxyFraction = 1;
		// defaults.receiveRecordThreshold = Long.MAX_VALUE;
		// defaults.receiverSignatureRequired = false;
		// defaults.sendRecordThreshold = Long.MAX_VALUE;
		// send create account transaction
		long shardToCreateIn = 0;
		long realmToCreateIn = 0;
		long startingBalance = 10;
		// let's create the account
		HederaTransactionResult createResult = myNewAccount.create(shardToCreateIn, realmToCreateIn,
				newAccountKey.getPublicKey(), newAccountKey.getKeyType(), startingBalance, defaults);
		// was it successful ?
		if (createResult.getPrecheckResult() == HederaPrecheckResult.OK) {
			// yes, get a receipt for the transaction
			HederaTransactionReceipt receipt = Utilities.getReceipt(myNewAccount.hederaTransactionID,
					myNewAccount.txQueryDefaults.node); // was that successful ?
			if (receipt.transactionStatus == HederaTransactionStatus.SUCCESS) {
				// yes, get the new account number from the receipt
				myNewAccount.accountNum = receipt.accountID.accountNum;
				// and print it out
				System.out.println(String.format("===>Your new account number is %d", myNewAccount.accountNum));
				// get balance
				myNewAccount.txQueryDefaults.payingAccountID = myNewAccount.getHederaAccountID();
				myNewAccount.txQueryDefaults.payingKeyPair = newAccountKey;
				myNewAccount.getBalance();
				HederaAccountID toAccountID = new HederaAccountID(nodeAccountShard, nodeAccountRealm, nodeAccountNum);
				myNewAccount.send(toAccountID, 1);
				
				Thread.sleep(100);
				myNewAccount.getBalance();
		    	// new file object
		    	HederaFile file = new HederaFile();
		    	// setup transaction/query defaults (durations, etc...)
		    	txQueryDefaults.fileWacl = new HederaCryptoKeyPair(KeyType.ED25519);
		    	file.txQueryDefaults = txQueryDefaults;

				boolean doCreate = false;
				boolean doGetInfo = false;
				boolean doAppend = false;
				boolean doUpdate = false;
				boolean doGetContents = false;
				boolean doDelete = false;

		    	//Test to create Files in HH
		    	doCreate = true; //OK
				doGetInfo = true; //OK
				doGetContents = true; //OK
				doAppend = false; //OK
				doUpdate = false; //OK
				doDelete = false; //OK
				
				// create a file
				if (doCreate) {
					String fileContents = "3-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789" + 
							"-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789";
					file =  FileCreate.create(file, fileContents.getBytes());
				}
		        if (file != null) {
		        	if (doGetInfo) {
		        		FileGetInfo.getInfo(file);
		        	}
		        	if (doAppend) {
			    		// append to a file
			            FileAppend.append(file,"Appended contents".getBytes());
			    		// get file contents
			        	if (doGetContents) {
			        		FileGetContents.getContents(file);
			        	}
		        	}
		        	if (doUpdate) {
		        		file = FileUpdate.update(file, 10, 2, "Updated contents".getBytes());
			    		// get file contents
			        	if (doGetContents) {
			        		FileGetContents.getContents(file);
			        	}
		        	}
		    		// get file contents
		        	if (doGetContents) {
		        		FileGetContents.getContents(file);
		        	}
		    		// get file info
		        	if (doGetInfo) {
		        		FileGetInfo.getInfo(file);
		        	}
		    		// delete the file
		        	if (doDelete) {
		        		FileDelete.delete(file);
		        	}
		        }
			}
		}
		// define the valid duration for the transactions (seconds, nanos)
	}
}