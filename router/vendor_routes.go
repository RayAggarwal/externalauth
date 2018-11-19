package router

const VendorRoutePrefix string = "/vendor"
func GetVendorRoutes() Routes {
	var routes = Routes{
		Route{
			"AddVendor",
			"POST",
			"",
			nil,
		},
	}
	return routes
}