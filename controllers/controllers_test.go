package controllers_test

import (
	"berbagi/config"
	"berbagi/controllers"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	config.InitDb()
	insertUser()
	insertAddress()
	insertDonor()
	insertChildren()
	insertFoundation()
	insertService()
	os.Exit(m.Run())
}

func TestRunner(t *testing.T) {
	t.Run("request gift", TestRequestGift)
	t.Run("request donation", TestRequestDonation)
	t.Run("request service", TestRequestService)
	t.Run("get all request", TestGetAllRequestList)
	t.Run("get type request", TestGetTypeRequestList)
	t.Run("get user id req", TestGetRequestByRecipientId)
	t.Run("make donation", TestMakeDonation)
	t.Run("make no request donation", TestNoRequestDonation)
	t.Run("make quick donation", TestQuickDonation)
	t.Run("get donations in cart", TestGetDonationInCart)
	t.Run("update donation cart", TestUpdateDonationInCart)
	t.Run("delete donation cart", TestDeleteDonationInCart)
	t.Run("checkout donation from cart", TestCheckoutDonation)
	t.Run("get donations list", TestGetDonationResolved)
	t.Run("pay donation", TestPayDonation)
	t.Run("nearby target", TestGetNearbyTarget)
	t.Run("delete request", TestDeleteRequest)

	t.Run("proficiency test", controllers.TestProficiency)
	t.Run("volunteer test", controllers.TestVolunteer)
	t.Run("services test", controllers.TestServices)
	t.Run("confirmation test", controllers.TestConfirmation)
	t.Run("completion test", controllers.TestCompletion)
	t.Run("certificate test", controllers.TestCertificate)

	controllers.InitCheckoutTest()
	t.Run("get checkout", controllers.Test_GetCheckoutByUserIdController)
	t.Run("add checkout", controllers.Test_AddCheckoutByUserIdController)
	t.Run("register user", controllers.Test_RegisterUserController)
	t.Run("login user", controllers.Test_LoginUserController)
	t.Run("get list products", controllers.Test_GetProductsController)
	t.Run("get product cart", controllers.Test_GetProductCartByUserIdController)
	t.Run("update product cart", controllers.Test_UpdateProductCartByUserIdController)
	t.Run("delete product cart", controllers.Test_DeleteProductCartByUserIdController)
	t.Run("get children gift", controllers.Test_GetGiftsController)
	t.Run("get pending payment", controllers.Test_GetPendingPaymentsController)
	t.Run("resolved payment", controllers.Test_AddPendingPaymentController)
}
