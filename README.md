# Basic TF testing repo

To manually run test locally:
  - Position inside test folder
  - Add AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY envirioment variables
  - Run `go test -v`
  
To run the TF files position on project root and run:
  - `terraform init`
  - `terraform apply`
 
 
If you wish to make changes make sure to add tests appropriately. All changes must be proposed through a Pull Request to the main branch. 
An automatic action will run all go tests inside the test folder. 
