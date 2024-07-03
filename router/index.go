package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	SetupAsset_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/asset_check"
	SetupAsset_count_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/asset_count_story"
	AssetsRoutes "github.com/puvadon-artmit/gofiber-template/api/assets"
	SetupAssets_count_StoreRoutes "github.com/puvadon-artmit/gofiber-template/api/assets_count_store"
	Assets_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/assets_story"
	authRoutes "github.com/puvadon-artmit/gofiber-template/api/auth"
	SetupAutoclik_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_check"
	SetupAutoclik_check_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_check_story"
	SetupAutoclik_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_count"
	SetupAutoclik_Count_ProductRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_posting"
	SetupAutoclik_count_StoreRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_store"
	SetupAutoclik_count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_count_story"
	SetupaAtoclik_counting_rightsRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_rights"
	SetupaAtoclik_counting_triggerRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_counting_trigger"
	SetupAutoclik_bin_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_data"
	SetupAutoclik_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_data"
	SetupAutoclik_fixed_assetRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset"
	SetupAsset_Fixed_Asset_CheckRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check"
	SetupAutoclik_Fixed_Asset_check_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_check_story"
	SetupAutoclik_Fixed_Asset_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_count"
	SetupAutoclik_Fixed_Asset_count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_count_story"
	SetupaAutoclik_Fixed_Asset_Counting_RightsRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_counting_rights"
	SetupAutoclik_Fixed_Asset_Photos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_photos"
	SetupAutoclik_fixed_asset_round_countRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_round_count"
	SetupAutoclik_fixed_asset_round_count_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_round_count_story"
	SetupAutoclik_Fixed_Asset_StoreRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_fixed_asset_store"
	SetupAutoclik_Photos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_photos_check"
	SetupAutoclik_Round_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_round_count"
	SetupAutoclik_Round_Count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_round_count_story"
	SetupAutoclik_Update_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/autoclik_update_story"
	branchRoutes "github.com/puvadon-artmit/gofiber-template/api/branch"
	SetupBranch_AutoclikRoutes "github.com/puvadon-artmit/gofiber-template/api/branch_autoclik"
	SetupBranch_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/branch_story"
	categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category"
	SetupCategory_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/category_story"
	SetupCount_AutoclikRoutes "github.com/puvadon-artmit/gofiber-template/api/count_autoclik"
	Setupcount_categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/count_category"
	Setupcount_main_categoryRoutes "github.com/puvadon-artmit/gofiber-template/api/count_main_category"
	Setupcounting_rightsRoutes "github.com/puvadon-artmit/gofiber-template/api/counting_rights"
	groundRouter "github.com/puvadon-artmit/gofiber-template/api/ground"
	Ground_storyRouter "github.com/puvadon-artmit/gofiber-template/api/ground_story"
	groupRoutes "github.com/puvadon-artmit/gofiber-template/api/group"
	SetupGroup_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/group_story"
	asset_countRoutes "github.com/puvadon-artmit/gofiber-template/api/inspection-assets"
	item_modelRoutes "github.com/puvadon-artmit/gofiber-template/api/item_model"
	SetupLocationRoutes "github.com/puvadon-artmit/gofiber-template/api/location"
	SetupLocation_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/location_story"
	SetupMain_branchRoutes "github.com/puvadon-artmit/gofiber-template/api/main-branch"
	SetupMain_Branch_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/main_branch_story"
	SetupMain_CategoryRoutes "github.com/puvadon-artmit/gofiber-template/api/main_category"
	SetupMain_Category_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/main_category_story"
	SetupMaliwan_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_check"
	SetupMaliwan_check_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_check_story"
	SetupMaliwan_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_count"
	SetupMaliwan_bin_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_count_store"
	SetupMaliwan_count_StoreRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_count_store"
	SetupaMaliwan_counting_rightsRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_rights"
	SetupaMaliwan_counting_triggerRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_counting_trigger"
	SetupMaliwan_counts_posting_groupRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_counts_posting_group"
	SetupMaliwan_counts_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_counts_story"
	SetupMaliwan_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_data"
	SetupMaliwan_Fixed_Asset_CheckRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_check"
	SetupMaliwan_Fixed_Asset_check_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_check_story"
	SetupMaliwan_Fixed_Asset_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_count"
	SetupMaliwan_Fixed_Asset_count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_count_story"
	SetupaMaliwan_Fixed_Asset_Counting_RightsRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_counting_rights"
	SetupMaliwan_Fixed_Asset_Photos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_photos"
	SetupMaliwan_fixed_asset_round_countRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_round_count"
	SetupMaliwan_fixed_asset_round_count_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_round_count_story"
	SetupMaliwan_Fixed_Asset_StoreRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_fixed_asset_store"
	SetupMaliwan_Photos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_photos_check"
	SetupMaliwan_Round_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_round_count"
	SetupMaliwan_Round_Count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_round_count_story"
	SetupMaliwan_Update_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/maliwan_update_story"
	manufacturerRoutes "github.com/puvadon-artmit/gofiber-template/api/manufacturer"
	SetupManufacturer_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/manufacturer_story"
	permissionRoutes "github.com/puvadon-artmit/gofiber-template/api/permission"
	SetupPermission_componentyRoutes "github.com/puvadon-artmit/gofiber-template/api/permissioncomponent"
	SetupPermissionGroupRoutes "github.com/puvadon-artmit/gofiber-template/api/permissiongroup"
	SetupPhotos_checkRoutes "github.com/puvadon-artmit/gofiber-template/api/photos_check"
	SetupPosting_groupsRoutes "github.com/puvadon-artmit/gofiber-template/api/posting_group"
	SetupRequest_update_dataRoutes "github.com/puvadon-artmit/gofiber-template/api/request_update_data"
	responsibleRoutes "github.com/puvadon-artmit/gofiber-template/api/responsible"
	SetupResponsible_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/responsible-story"
	roleRoutes "github.com/puvadon-artmit/gofiber-template/api/role"
	SetupRound_CountRoutes "github.com/puvadon-artmit/gofiber-template/api/round_count"
	SetupRound_Count_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/round_count_story"
	SetupScan_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/scan_story"
	SetupSignatureRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation"
	SetupSignature_AutoclikeRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik"
	SetupSignature_Autoclike_Fixed_AssetRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-autoclik-fixed-asset"
	SetupSignature_MaliwanRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-maliwan"
	SetupSignature_Maliwan_Fixed_AssetRoutes "github.com/puvadon-artmit/gofiber-template/api/sign-confirmation-maliwan-fixed-asset"
	statusRoutes "github.com/puvadon-artmit/gofiber-template/api/status"
	SetupStatus_StoryRoutes "github.com/puvadon-artmit/gofiber-template/api/status_story"
	storyRoutes "github.com/puvadon-artmit/gofiber-template/api/story"
	typeRoutes "github.com/puvadon-artmit/gofiber-template/api/type_things"
	SetupType_things_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/type_things_story"
	typeplanRoutes "github.com/puvadon-artmit/gofiber-template/api/typeplan"
	typeplan_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/typeplan_story"
	SetupUser_storyRoutes "github.com/puvadon-artmit/gofiber-template/api/user_story"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())

	// api := app.Group("api", middleware.NewECSLoggerMiddleWare)
	// ใช้ตอน deploy

	// app.Use(fiberzap.New(fiberzap.Config{
	// 	Logger: log.Log,
	// }))
	// authRoutes.SetupAuthRoutes(api)
	authRoutes.SetupAuthRoutes(api)
	roleRoutes.SetupRoleRoutes(api)
	permissionRoutes.SetupPermissionRoutes(api)
	categoryRoutes.SetupCategoryRoutes(api)
	branchRoutes.SetupBranchRoutes(api)
	groupRoutes.SetupGroupRoutes(api)
	statusRoutes.SetupStatusRoutes(api)
	typeRoutes.SetupTypeRoutes(api)
	manufacturerRoutes.SetupManufacturerRoutes(api)
	item_modelRoutes.SetupItem_modelRoutes(api)
	storyRoutes.SetupStoryRoutes(api)
	groundRouter.SetupGroundRoutes(api)
	Ground_storyRouter.SetupGround_storyRoutes(api)
	responsibleRoutes.SetupResponsibleRoutes(api)
	AssetsRoutes.SetupAssetsRoutes(api)
	Assets_StoryRoutes.SetupAssets_StoryRoutes(api)
	typeplanRoutes.SetupTypeplanRoutes(api)
	typeplan_storyRoutes.SetupTypeplan_StoryRoutes(api)
	asset_countRoutes.SetupAsset_countRoutes(api)
	SetupAsset_checkRoutes.SetupAsset_checkRoutes(api)
	SetupLocationRoutes.SetupLocationRoutes(api)
	SetupLocation_storyRoutes.SetupLocation_storyRoutes(api)
	SetupSignatureRoutes.SetupSignatureRoutes(api)
	SetupMain_branchRoutes.SetupMain_branchRoutes(api)
	SetupBranch_StoryRoutes.SetupBranch_StoryRoutes(api)
	SetupMain_Branch_StoryRoutes.SetupMain_Branch_StoryRoutes(api)
	SetupGroup_StoryRoutes.SetupGroup_StoryRoutes(api)
	SetupResponsible_StoryRoutes.SetupResponsible_StoryRoutes(api)
	SetupAutoclik_dataRoutes.SetupAutoclik_dataRoutes(api)
	SetupMaliwan_dataRoutes.SetupMaliwan_dataRoutes(api)
	SetupRound_CountRoutes.SetupRound_CountRoutes(api)
	SetupRound_Count_StoryRoutes.SetupRound_Count_StoryRoutes(api)
	SetupMain_CategoryRoutes.SetupMain_CategoryRoutes(api)
	SetupBranch_AutoclikRoutes.SetupBranch_AutoclikRoutes(api)
	SetupCount_AutoclikRoutes.SetupCount_AutoclikRoutes(api)
	SetupPhotos_checkRoutes.SetupPhotos_checkRoutes(api)
	Setupcount_categoryRoutes.Setupcount_categoryRoutes(api)
	Setupcount_main_categoryRoutes.Setupcount_main_categoryRoutes(api)
	Setupcount_categoryRoutes.Setupcount_categoryRoutes(api)
	Setupcount_main_categoryRoutes.Setupcount_main_categoryRoutes(api)
	SetupPermission_componentyRoutes.SetupPermission_componentyRoutes(api)
	SetupPermissionGroupRoutes.SetupPermissionGroupRoutes(api)
	SetupCategory_StoryRoutes.SetupCategory_StoryRoutes(api)
	Setupcounting_rightsRoutes.Setupcounting_rightsRoutes(api)
	SetupManufacturer_storyRoutes.SetupManufacturer_storyRoutes(api)
	SetupAsset_count_storyRoutes.SetupAsset_count_storyRoutes(api)
	SetupScan_storyRoutes.SetupScan_storyRoutes(api)
	SetupMain_Category_storyRoutes.SetupMain_Category_storyRoutes(api)
	SetupStatus_StoryRoutes.SetupStatus_StoryRoutes(api)
	SetupType_things_storyRoutes.SetupType_things_storyRoutes(api)
	SetupAutoclik_CountRoutes.SetupAutoclik_CountRoutes(api)
	SetupAutoclik_count_StoryRoutes.SetupAutoclik_count_StoryRoutes(api)
	SetupAutoclik_Round_CountRoutes.SetupAutoclik_Round_CountRoutes(api)
	SetupAutoclik_Round_Count_StoryRoutes.SetupAutoclik_Round_Count_StoryRoutes(api)
	SetupAutoclik_checkRoutes.SetupAutoclik_checkRoutes(api)
	SetupaAtoclik_counting_rightsRoutes.SetupaAtoclik_counting_rightsRoutes(api)
	SetupAutoclik_Count_ProductRoutes.SetupAutoclik_Count_ProductRoutes(api)
	SetupAutoclik_Photos_checkRoutes.SetupAutoclik_Photos_checkRoutes(api)
	SetupAutoclik_bin_dataRoutes.SetupAutoclik_bin_dataRoutes(api)
	SetupAutoclik_check_StoryRoutes.SetupAutoclik_check_StoryRoutes(api)
	SetupSignature_AutoclikeRoutes.SetupSignature_AutoclikeRoutes(api)
	SetupAutoclik_Update_StoryRoutes.SetupAutoclik_Update_StoryRoutes(api)
	SetupAutoclik_fixed_assetRoutes.SetupAutoclik_fixed_assetRoutes(api)
	SetupMaliwan_Update_StoryRoutes.SetupMaliwan_Update_StoryRoutes(api)
	SetupUser_storyRoutes.SetupUser_storyRoutes(api)
	SetupMaliwan_CountRoutes.SetupMaliwan_CountRoutes(api)
	SetupMaliwan_counts_storyRoutes.SetupMaliwan_counts_storyRoutes(api)
	SetupMaliwan_Round_CountRoutes.SetupMaliwan_Round_CountRoutes(api)
	SetupMaliwan_counts_posting_groupRoutes.SetupMaliwan_counts_posting_groupRoutes(api)
	SetupMaliwan_Round_Count_StoryRoutes.SetupMaliwan_Round_Count_StoryRoutes(api)
	SetupaMaliwan_counting_rightsRoutes.SetupaMaliwan_counting_rightsRoutes(api)
	SetupMaliwan_checkRoutes.SetupMaliwan_checkRoutes(api)
	SetupMaliwan_Photos_checkRoutes.SetupMaliwan_Photos_checkRoutes(api)
	SetupSignature_MaliwanRoutes.SetupSignature_MaliwanRoutes(api)
	SetupMaliwan_check_StoryRoutes.SetupMaliwan_check_StoryRoutes(api)
	SetupaAtoclik_counting_triggerRoutes.SetupaAtoclik_counting_triggerRoutes(api)
	SetupPosting_groupsRoutes.SetupPosting_groupsRoutes(api)
	SetupAutoclik_count_StoreRoutes.SetupAutoclik_count_StoreRoutes(api)
	SetupaMaliwan_counting_triggerRoutes.SetupaMaliwan_counting_triggerRoutes(api)
	SetupMaliwan_count_StoreRoutes.SetupMaliwan_count_StoreRoutes(api)
	SetupMaliwan_bin_dataRoutes.SetupMaliwan_bin_dataRoutes(api)
	SetupAutoclik_Fixed_Asset_CountRoutes.SetupAutoclik_Fixed_Asset_CountRoutes(api)
	SetupAutoclik_Fixed_Asset_StoreRoutes.SetupAutoclik_Fixed_Asset_StoreRoutes(api)
	SetupaAutoclik_Fixed_Asset_Counting_RightsRoutes.SetupaAutoclik_Fixed_Asset_Counting_RightsRoutes(api)
	SetupAutoclik_fixed_asset_round_countRoutes.SetupAutoclik_fixed_asset_round_countRoutes(api)
	SetupAutoclik_fixed_asset_round_count_storyRoutes.SetupAutoclik_fixed_asset_round_count_storyRoutes(api)
	SetupAutoclik_Fixed_Asset_Photos_checkRoutes.SetupAutoclik_Fixed_Asset_Photos_checkRoutes(api)
	SetupAsset_Fixed_Asset_CheckRoutes.SetupAsset_Fixed_Asset_CheckRoutes(api)
	SetupAutoclik_Fixed_Asset_count_StoryRoutes.SetupAutoclik_Fixed_Asset_count_StoryRoutes(api)
	SetupAutoclik_Fixed_Asset_check_StoryRoutes.SetupAutoclik_Fixed_Asset_check_StoryRoutes(api)
	SetupSignature_Autoclike_Fixed_AssetRoutes.SetupSignature_Autoclike_Fixed_AssetRoutes(api)
	SetupAssets_count_StoreRoutes.SetupAssets_count_StoreRoutes(api)
	SetupMaliwan_Fixed_Asset_CountRoutes.SetupMaliwan_Fixed_Asset_CountRoutes(api)
	SetupMaliwan_Fixed_Asset_count_StoryRoutes.SetupMaliwan_Fixed_Asset_count_StoryRoutes(api)
	SetupMaliwan_fixed_asset_round_countRoutes.SetupMaliwan_fixed_asset_round_countRoutes(api)
	SetupMaliwan_fixed_asset_round_count_storyRoutes.SetupMaliwan_fixed_asset_round_count_storyRoutes(api)
	SetupaMaliwan_Fixed_Asset_Counting_RightsRoutes.SetupaMaliwan_Fixed_Asset_Counting_RightsRoutes(api)
	SetupMaliwan_Fixed_Asset_Photos_checkRoutes.SetupMaliwan_Fixed_Asset_Photos_checkRoutes(api)
	SetupMaliwan_Fixed_Asset_CheckRoutes.SetupMaliwan_Fixed_Asset_CheckRoutes(api)
	SetupSignature_Maliwan_Fixed_AssetRoutes.SetupSignature_Maliwan_Fixed_AssetRoutes(api)
	SetupMaliwan_Fixed_Asset_StoreRoutes.SetupMaliwan_Fixed_Asset_StoreRoutes(api)
	SetupMaliwan_Fixed_Asset_check_StoryRoutes.SetupMaliwan_Fixed_Asset_check_StoryRoutes(api)
	SetupRequest_update_dataRoutes.SetupRequest_update_dataRoutes(api)
}
