package routesname

const BasePath = "/api/v1"

const UserBasePath = BasePath + "/users"
const UserLogin = UserBasePath + "/login"
const UserRegister = UserBasePath + "/register"

const CollectorBasePath = BasePath + "/collectors"
const CollectorLogin = CollectorBasePath + "/login"
const CollectorRegister = CollectorBasePath + "/register"

const PickupBasePath = BasePath + "/pickups"
const PickupById = PickupBasePath + "/:id"

const TransactionBasePath = BasePath + "/transactions"
const TransactionById = TransactionBasePath + "/:id"

const ChatBasePath = BasePath + "/chats"
const ChatMessage = BasePath + "/message"

const MediaUpload = BasePath + "/media"