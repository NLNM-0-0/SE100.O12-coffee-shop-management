/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `Category`;
CREATE TABLE `Category` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` varchar(200) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci DEFAULT '',
  `amountProduct` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `CategoryFood`;
CREATE TABLE `CategoryFood` (
  `foodId` varchar(12) NOT NULL,
  `categoryId` varchar(12) NOT NULL,
  PRIMARY KEY (`foodId`,`categoryId`),
  KEY `categoryId` (`categoryId`),
  CONSTRAINT `CategoryFood_ibfk_1` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `CategoryFood_ibfk_2` FOREIGN KEY (`categoryId`) REFERENCES `Category` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Customer`;
CREATE TABLE `Customer` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `point` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNote`;
CREATE TABLE `ExportNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `reason` enum('Damaged','OutOfDate') DEFAULT NULL,
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `ExportNote_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ExportNoteDetail`;
CREATE TABLE `ExportNoteDetail` (
  `exportNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountExport` float DEFAULT '0',
  `unitTypeName` text NOT NULL,
  PRIMARY KEY (`exportNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `ExportNoteDetail_ibfk_1` FOREIGN KEY (`exportNoteId`) REFERENCES `ExportNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ExportNoteDetail_ibfk_2` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Feature`;
CREATE TABLE `Feature` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `description` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `groupName` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Food`;
CREATE TABLE `Food` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text NOT NULL,
  `cookingGuide` text NOT NULL,
  `image` text NOT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNote`;
CREATE TABLE `ImportNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `supplierId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `totalPrice` int DEFAULT '0',
  `status` enum('InProgress','Done','Cancel') DEFAULT 'InProgress',
  `createdBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `closedBy` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `closedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `closedBy` (`closedBy`),
  KEY `supplierId` (`supplierId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `ImportNote_ibfk_1` FOREIGN KEY (`closedBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_2` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNote_ibfk_3` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ImportNoteDetail`;
CREATE TABLE `ImportNoteDetail` (
  `importNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` int NOT NULL,
  `amountImport` float NOT NULL,
  `totalUnit` int NOT NULL,
  `unitTypeName` text NOT NULL,
  PRIMARY KEY (`importNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `ImportNoteDetail_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `ImportNoteDetail_ibfk_2` FOREIGN KEY (`importNoteId`) REFERENCES `ImportNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Ingredient`;
CREATE TABLE `Ingredient` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `amount` float DEFAULT '0',
  `unitTypeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `unitTypeId` (`unitTypeId`),
  CONSTRAINT `Ingredient_ibfk_1` FOREIGN KEY (`unitTypeId`) REFERENCES `UnitType` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNote`;
CREATE TABLE `InventoryCheckNote` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(12) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `InventoryCheckNote_ibfk_1` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InventoryCheckNoteDetail`;
CREATE TABLE `InventoryCheckNoteDetail` (
  `inventoryCheckNoteId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) NOT NULL,
  `initial` float NOT NULL,
  `difference` float NOT NULL,
  `final` float NOT NULL,
  PRIMARY KEY (`inventoryCheckNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `InventoryCheckNoteDetail_ibfk_2` FOREIGN KEY (`inventoryCheckNoteId`) REFERENCES `InventoryCheckNote` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Invoice`;
CREATE TABLE `Invoice` (
  `id` varchar(13) NOT NULL,
  `customerId` varchar(13) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `totalPrice` int NOT NULL,
  `totalCost` int DEFAULT NULL,
  `amountReceived` int NOT NULL,
  `amountPriceUsePoint` int NOT NULL,
  `pointUse` int NOT NULL,
  `pointReceive` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(13) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `customerId` (`customerId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `Invoice_ibfk_2` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `Invoice_ibfk_3` FOREIGN KEY (`customerId`) REFERENCES `Customer` (`id`) ON DELETE SET DEFAULT ON UPDATE SET DEFAULT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `InvoiceDetail`;
CREATE TABLE `InvoiceDetail` (
  `invoiceId` varchar(13) NOT NULL,
  `foodId` varchar(13) NOT NULL,
  `sizeName` text NOT NULL,
  `amount` int NOT NULL,
  `unitPrice` int NOT NULL,
  `description` text NOT NULL,
  `toppings` json NOT NULL,
  KEY `invoiceId` (`invoiceId`) USING BTREE,
  KEY `foodId` (`foodId`),
  CONSTRAINT `InvoiceDetail_ibfk_1` FOREIGN KEY (`invoiceId`) REFERENCES `Invoice` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `InvoiceDetail_ibfk_2` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `MUser`;
CREATE TABLE `MUser` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `image` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `roleId` (`roleId`),
  CONSTRAINT `MUser_ibfk_1` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Recipe`;
CREATE TABLE `Recipe` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RecipeDetail`;
CREATE TABLE `RecipeDetail` (
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amountNeed` float NOT NULL,
  PRIMARY KEY (`recipeId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `RecipeDetail_ibfk_1` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `RecipeDetail_ibfk_2` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Role`;
CREATE TABLE `Role` (
  `id` varchar(13) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `RoleFeature`;
CREATE TABLE `RoleFeature` (
  `roleId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `featureId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`roleId`,`featureId`),
  KEY `featureId` (`featureId`),
  CONSTRAINT `RoleFeature_ibfk_1` FOREIGN KEY (`featureId`) REFERENCES `Feature` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `RoleFeature_ibfk_2` FOREIGN KEY (`roleId`) REFERENCES `Role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `ShopGeneral`;
CREATE TABLE `ShopGeneral` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `email` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `phone` text NOT NULL,
  `address` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `wifiPass` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `accumulatePointPercent` float NOT NULL,
  `usePointPercent` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SizeFood`;
CREATE TABLE `SizeFood` (
  `foodId` varchar(12) NOT NULL,
  `sizeId` varchar(12) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `cost` int NOT NULL,
  `price` int NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`foodId`,`sizeId`),
  KEY `recipeId` (`recipeId`),
  CONSTRAINT `SizeFood_ibfk_1` FOREIGN KEY (`foodId`) REFERENCES `Food` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SizeFood_ibfk_2` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `StockChangeHistory`;
CREATE TABLE `StockChangeHistory` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `ingredientId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` float NOT NULL DEFAULT '0',
  `amountLeft` float NOT NULL DEFAULT '0',
  `type` enum('Sell','Import','Export','Modify') NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`,`ingredientId`),
  KEY `supplierId` (`ingredientId`),
  CONSTRAINT `StockChangeHistory_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `StockReport`;
CREATE TABLE `StockReport` (
  `id` varchar(12) NOT NULL,
  `timeFrom` timestamp NOT NULL,
  `timeTo` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `StockReportDetail`;
CREATE TABLE `StockReportDetail` (
  `reportId` varchar(12) NOT NULL,
  `ingredientId` varchar(12) NOT NULL,
  `initial` float NOT NULL,
  `sell` float NOT NULL,
  `import` float NOT NULL,
  `export` float NOT NULL,
  `modify` float NOT NULL,
  `final` float NOT NULL,
  PRIMARY KEY (`reportId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `StockReportDetail_ibfk_1` FOREIGN KEY (`reportId`) REFERENCES `StockReport` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `StockReportDetail_ibfk_2` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Supplier`;
CREATE TABLE `Supplier` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` text NOT NULL,
  `email` text NOT NULL,
  `phone` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `debt` int DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebt`;
CREATE TABLE `SupplierDebt` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `supplierId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `amount` int NOT NULL,
  `amountLeft` int NOT NULL,
  `type` enum('Debt','Pay') NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `createdBy` varchar(9) NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `supplierId` (`supplierId`),
  KEY `createdBy` (`createdBy`),
  CONSTRAINT `SupplierDebt_ibfk_1` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SupplierDebt_ibfk_2` FOREIGN KEY (`createdBy`) REFERENCES `MUser` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebtReport`;
CREATE TABLE `SupplierDebtReport` (
  `id` varchar(12) NOT NULL,
  `timeFrom` timestamp NOT NULL,
  `timeTo` timestamp NOT NULL,
  `initial` int NOT NULL,
  `debt` int NOT NULL,
  `pay` int NOT NULL,
  `final` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `SupplierDebtReportDetail`;
CREATE TABLE `SupplierDebtReportDetail` (
  `reportId` varchar(12) NOT NULL,
  `supplierId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `debt` int NOT NULL,
  `pay` int NOT NULL,
  `final` int NOT NULL,
  PRIMARY KEY (`reportId`,`supplierId`),
  KEY `supplierId` (`supplierId`),
  CONSTRAINT `SupplierDebtReportDetail_ibfk_1` FOREIGN KEY (`supplierId`) REFERENCES `Supplier` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `SupplierDebtReportDetail_ibfk_2` FOREIGN KEY (`reportId`) REFERENCES `SupplierDebtReport` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Topping`;
CREATE TABLE `Topping` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `description` text NOT NULL,
  `cookingGuide` text NOT NULL,
  `recipeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  `cost` int NOT NULL,
  `price` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `recipeId` (`recipeId`),
  CONSTRAINT `Topping_ibfk_1` FOREIGN KEY (`recipeId`) REFERENCES `Recipe` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `UnitType`;
CREATE TABLE `UnitType` (
  `id` varchar(13) NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `value` int DEFAULT NULL,
  `measureType` enum('Weight','Volume','Unit') DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('8y1Wu_KIR', 'Sinh Tố', '', 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('BBsWulFSg', 'Nước ép', '', 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('buQgXlKIg', 'Cà phê', '', 2);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('ORLgXlKIg', 'Trà', '', 2),
('rmtWX_FIg', 'Sữa chua', '', 0),
('SYAnu_KIR', 'Trà trái cây', '', 1),
('w6yMulFIR', 'Bánh ngọt', '', 0);

INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('9Ip99_FSR', 'buQgXlKIg');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('nMwZjlKIR', 'buQgXlKIg');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('0yEY9_KSR', 'ORLgXlKIg');
INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('oiNS9_FSg', 'ORLgXlKIg'),
('oiNS9_FSg', 'SYAnu_KIR');

INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('0VHUIXFIR', 'Nguyễn Lê Ngọc Mai', 'mai@gmail.com', '0912393894', 350);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('pD0cSuKIR', 'Nguyễn Kim Anh Thư', 'anhthu@gmail.com', '0983219471', 4750);


INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('dsIKBXKIg', 'Damaged', 'g3W21A7SR', '2023-11-10 05:03:58');
INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('ofuFfXKSR', 'OutOfDate', 'g3W21A7SR', '2023-11-12 05:04:18');


INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`, `unitTypeName`) VALUES
('dsIKBXKIg', 'NVLTuiTraD', 10, 'đơn vị');
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`, `unitTypeName`) VALUES
('ofuFfXKSR', 'NVLDao', 10, 'đơn vị');
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`, `unitTypeName`) VALUES
('ofuFfXKSR', 'NVLTrung', 3, 'đơn vị');

INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_CREATE', 'Tạo danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_UP_INFO', 'Chỉnh sửa thông tin danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CAT_VIEW', 'Xem danh mục', 'Danh mục');
INSERT INTO `Feature` (`id`, `description`, `groupName`) VALUES
('CUS_CREATE', 'Tạo khách hàng', 'Khách hàng'),
('CUS_UP_INFO', 'Chỉnh sửa thông tin khách hàng', 'Khách hàng'),
('CUS_VIEW', 'Xem khách hàng', 'Khách hàng'),
('EXP_CREATE', 'Tạo phiếu xuất', 'Phiếu xuất'),
('EXP_VIEW', 'Xem phiếu xuất', 'Phiếu xuất'),
('FOD_CREATE', 'Tạo mặt hàng', 'Mặt hàng'),
('FOD_UP_INFO', 'Chỉnh sửa thông tin mặt hàng', 'Mặt hàng'),
('FOD_UP_STATE', 'Chỉnh sửa trạng thái mặt hàng', 'Mặt hàng'),
('FOD_VIEW', 'Xem mặt hàng', 'Mặt hàng'),
('ICN_CREATE', 'Tạo phiếu kiểm kho', 'Phiếu kiểm kho'),
('ICN_VIEW', 'Xem phiếu kiểm kho', 'Phiếu kiểm kho'),
('IMP_CREATE', 'Tạo phiếu nhập', 'Phiếu nhập'),
('IMP_UP_STATE', 'Chỉnh sửa trạng thái phiếu nhập', 'Phiếu nhập'),
('IMP_VIEW', 'Xem phiếu nhập', 'Phiếu nhập'),
('ING_CREATE', 'Tạo nguyên liệu', 'Nguyên liệu'),
('ING_UP', 'Chỉnh sửa thông tin nguyên liệu', 'Nguyên liệu'),
('ING_VIEW', 'Xem nguyên liệu', 'Nguyên liệu'),
('INV_CREATE', 'Tạo hóa đơn', 'Hóa đơn'),
('INV_VIEW', 'Xem hóa đơn', 'Hóa đơn'),
('RPT_DEBT', 'Xem báo cáo nợ', 'Báo cáo'),
('RPT_SALE', 'Xem báo cáo mặt hàng', 'Báo cáo'),
('RPT_STOCK', 'Xem báo cáo tồn kho', 'Báo cáo'),
('SUP_CREATE', 'Tạo nhà cung cấp', 'Nhà cung cấp'),
('SUP_PAY', 'Trả nợ nhà cung cấp', 'Nhà cung cấp'),
('SUP_UP_INFO', 'Chỉnh sửa thông tin nhà cung cấp', 'Nhà cung cấp'),
('SUP_VIEW', 'Xem nhà cung cấp', 'Nhà cung cấp'),
('TOP_CREATE', 'Tạo topping', 'Topping'),
('TOP_UP_INFO', 'Chỉnh sửa thông tin topping', 'Topping'),
('TOP_UP_STATE', 'Chỉnh sửa trạng thái topping', 'Topping'),
('TOP_VIEW', 'Xem topping', 'Topping'),
('USE_UP_INFO', 'Chỉnh sửa thông tin người dùng', 'Nhân viên'),
('USE_VIEW', 'Xem người dùng', 'Nhân viên');

INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `image`, `isActive`) VALUES
('0yEY9_KSR', 'Trà ô long', '', '', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Food%2Ftrà sữa ô long nướngahfY9lKSR?alt=media', 1);
INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `image`, `isActive`) VALUES
('9Ip99_FSR', 'Cà phê đen', '', '', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Food%2Fcà phê đenlgfqr_FSg?alt=media', 1);
INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `image`, `isActive`) VALUES
('nMwZjlKIR', 'Bạc Sỉu', '', '', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Food%2Fbạc sỉu9ui2ClKSR?alt=media', 1);
INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `image`, `isActive`) VALUES
('oiNS9_FSg', 'Trà Đào Cam Sả', '', '', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Food%2Ftrà đào cam xảG9MI9_FSR?alt=media', 1);

INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('flMvfuFIg', 'NCCHorecavn', 5500000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-11-04 05:02:49', '2023-11-04 23:55:23');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('J5CSfXKSR', 'NCCAbby', 4900000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-11-03 05:02:09', '2023-11-03 16:05:42');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('kq5VfXKIg', 'NCCBeemart', 3480000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-11-02 05:01:19', '2023-11-02 22:00:53');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('y9zOBuFSR', 'NCCTrumNL', 1800000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2023-11-05 05:03:19', '2023-11-05 12:12:45');

INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('flMvfuFIg', 'NVLDao', 5000, 100, 500000, 'đơn vị');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('flMvfuFIg', 'NVLDen', 25000, 100, 2500000, 'kg');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('flMvfuFIg', 'NVLTrang', 25000, 100, 2500000, 'kg');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('J5CSfXKSR', 'NVLOLong', 1200000, 3, 3600000, 'kg'),
('J5CSfXKSR', 'NVLSua', 20000, 50, 1000000, 'l'),
('J5CSfXKSR', 'NVLTuiTraD', 1500, 200, 300000, 'đơn vị'),
('kq5VfXKIg', 'NVLBo', 1000, 3000, 3000000, 'g'),
('kq5VfXKIg', 'NVLBotNang', 30, 4000, 120000, 'g'),
('kq5VfXKIg', 'NVLBotSua', 60000, 1, 60000, 'kg'),
('kq5VfXKIg', 'NVLNesCafe', 3000, 100, 300000, 'đơn vị'),
('y9zOBuFSR', 'NVLDuong', 30000, 50, 1500000, 'kg'),
('y9zOBuFSR', 'NVLTrung', 3000, 100, 300000, 'đơn vị');

INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('NVLBo', 'Bơ', 2000, 'g', 1000);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('NVLBotNang', 'Bột năng', 4000, 'g', 30);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('NVLBotSua', 'Bột Sữa', 2.984, 'kg', 60000);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('NVLDao', 'Đào Miếng', 70, 'dv', 5000),
('NVLDen', 'Trân châu đen', 99.6, 'kg', 25000),
('NVLDuong', 'Đường', 57.32, 'kg', 30000),
('NVLNesCafe', 'Gói NesCafe', 89, 'dv', 3000),
('NVLOLong', 'Trà ô long', 2.84, 'kg', 1200000),
('NVLSua', 'Sữa', 46.4, 'l', 20000),
('NVLTrang', 'Trân châu trắng', 99.5, 'kg', 25000),
('NVLTrung', 'Trứng', 93, 'dv', 3000),
('NVLTuiTraD', 'Trà Đào Cozy', 185, 'dv', 1500);

INSERT INTO `InventoryCheckNote` (`id`, `createdAt`, `createdBy`) VALUES
('1cx_BuFIR', '2023-11-21 05:12:20', 'g3W21A7SR');
INSERT INTO `InventoryCheckNote` (`id`, `createdAt`, `createdBy`) VALUES
('RiWlBuKIR', '2023-11-20 05:04:18', 'g3W21A7SR');
INSERT INTO `InventoryCheckNote` (`id`, `createdAt`, `createdBy`) VALUES
('uDK_BuKIR', '2023-11-25 05:12:13', 'g3W21A7SR');

INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `ingredientId`, `initial`, `difference`, `final`) VALUES
('1cx_BuFIR', 'NVLBo', 3000, -1000, 2000);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `ingredientId`, `initial`, `difference`, `final`) VALUES
('RiWlBuKIR', 'NVLBotSua', 1, 1, 2);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `ingredientId`, `initial`, `difference`, `final`) VALUES
('RiWlBuKIR', 'NVLDao', 90, -10, 80);
INSERT INTO `InventoryCheckNoteDetail` (`inventoryCheckNoteId`, `ingredientId`, `initial`, `difference`, `final`) VALUES
('uDK_BuKIR', 'NVLBotSua', 2, 1, 3),
('uDK_BuKIR', 'NVLDuong', 50, 10, 60);

INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalCost`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdAt`, `createdBy`) VALUES
('5NzPEuFIg', '0VHUIXFIR', 105000, 35000, 105000, 0, 0, 10500, '2023-12-08 01:12:13', 'g3W21A7SR');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalCost`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdAt`, `createdBy`) VALUES
('65uBPuFIg', 'pD0cSuKIR', 105000, 10000, 105000, 0, 0, 10500, '2023-12-04 15:12:13', 'g3W21A7SR');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalCost`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdAt`, `createdBy`) VALUES
('awrYPXKSg', 'pD0cSuKIR', 90000, 40000, 90000, 0, 0, 9000, '2023-12-06 23:12:13', 'g3W21A7SR');
INSERT INTO `Invoice` (`id`, `customerId`, `totalPrice`, `totalCost`, `amountReceived`, `amountPriceUsePoint`, `pointUse`, `pointReceive`, `createdAt`, `createdBy`) VALUES
('DJNPPXKIg', '0VHUIXFIR', 105000, 44000, 105000, 0, 0, 10500, '2023-12-09 04:12:13', 'g3W21A7SR'),
('FAoBEuFIg', NULL, 20000, 12000, 20000, 0, 0, 0, '2023-12-03 05:12:13', 'g3W21A7SR'),
('HsJLPXKSg', 'pD0cSuKIR', 130000, 30000, 130000, 0, 0, 13000, '2023-12-05 17:12:13', 'g3W21A7SR'),
('uX3QEXKIR', '0VHUIXFIR', 105000, 44000, 105000, 0, 0, 10500, '2023-12-17 04:56:44', 'g3W21A7SR'),
('wQx_EuFSR', '0VHUIXFIR', 35000, 10000, 3500, 31500, 31500, 350, '2023-12-30 08:56:44', 'g3W21A7SR'),
('YhDlEXFSg', NULL, 105000, 10000, 105000, 0, 0, 0, '2023-12-30 05:56:44', 'g3W21A7SR'),
('Zf8wEXFSg', 'pD0cSuKIR', 80000, 12000, 47500, 32500, 32500, 4750, '2023-12-11 04:56:44', 'g3W21A7SR');

INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('FAoBEuFIg', 'nMwZjlKIR', 'S', 1, 20000, '', 'null');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('65uBPuFIg', 'oiNS9_FSg', 'S', 1, 45000, '', 'null');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('65uBPuFIg', '0yEY9_KSR', 'S', 1, 45000, '', 'null');
INSERT INTO `InvoiceDetail` (`invoiceId`, `foodId`, `sizeName`, `amount`, `unitPrice`, `description`, `toppings`) VALUES
('65uBPuFIg', '9Ip99_FSR', 'S', 1, 15000, '', 'null'),
('HsJLPXKSg', '0yEY9_KSR', 'S', 1, 55000, '', '[{\"id\": \"BTzlu_FSR\", \"name\": \"Pudding trứng\", \"price\": 5000}, {\"id\": \"Rnqfu_FSR\", \"name\": \"Trân châu trắng\", \"price\": 5000}]'),
('HsJLPXKSg', '9Ip99_FSR', 'S', 1, 30000, '', '[{\"id\": \"Rnqfu_FSR\", \"name\": \"Trân châu trắng\", \"price\": 5000}, {\"id\": \"BTzlu_FSR\", \"name\": \"Pudding trứng\", \"price\": 5000}, {\"id\": \"hEMEu_KIg\", \"name\": \"Trân châu đen\", \"price\": 5000}]'),
('HsJLPXKSg', 'oiNS9_FSg', 'S', 1, 45000, '', 'null'),
('awrYPXKSg', 'nMwZjlKIR', 'S', 1, 20000, '', 'null'),
('awrYPXKSg', '9Ip99_FSR', 'S', 1, 15000, '', 'null'),
('awrYPXKSg', '0yEY9_KSR', 'S', 1, 55000, '', '[{\"id\": \"Rnqfu_FSR\", \"name\": \"Trân châu trắng\", \"price\": 5000}, {\"id\": \"hEMEu_KIg\", \"name\": \"Trân châu đen\", \"price\": 5000}]'),
('5NzPEuFIg', '9Ip99_FSR', 'S', 1, 15000, '', 'null'),
('5NzPEuFIg', 'oiNS9_FSg', 'S', 1, 45000, '', 'null'),
('5NzPEuFIg', '0yEY9_KSR', 'S', 1, 45000, '', 'null'),
('DJNPPXKIg', 'oiNS9_FSg', 'S', 1, 45000, '', 'null'),
('DJNPPXKIg', '0yEY9_KSR', 'S', 1, 60000, '', '[{\"id\": \"hEMEu_KIg\", \"name\": \"Trân châu đen\", \"price\": 5000}, {\"id\": \"Rnqfu_FSR\", \"name\": \"Trân châu trắng\", \"price\": 5000}, {\"id\": \"BTzlu_FSR\", \"name\": \"Pudding trứng\", \"price\": 5000}]'),
('Zf8wEXFSg', '9Ip99_FSR', 'S', 1, 15000, '', 'null'),
('Zf8wEXFSg', '0yEY9_KSR', 'S', 1, 45000, '', 'null'),
('Zf8wEXFSg', 'nMwZjlKIR', 'S', 1, 20000, '', 'null'),
('uX3QEXKIR', 'oiNS9_FSg', 'S', 1, 45000, '', 'null'),
('uX3QEXKIR', '0yEY9_KSR', 'S', 1, 60000, '', '[{\"id\": \"hEMEu_KIg\", \"name\": \"Trân châu đen\", \"price\": 5000}, {\"id\": \"BTzlu_FSR\", \"name\": \"Pudding trứng\", \"price\": 5000}, {\"id\": \"Rnqfu_FSR\", \"name\": \"Trân châu trắng\", \"price\": 5000}]'),
('YhDlEXFSg', '0yEY9_KSR', 'S', 2, 45000, '', 'null'),
('YhDlEXFSg', '9Ip99_FSR', 'S', 1, 15000, '', 'null'),
('wQx_EuFSR', 'nMwZjlKIR', 'S', 1, 20000, '', 'null'),
('wQx_EuFSR', '9Ip99_FSR', 'S', 1, 15000, '', 'null');

INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('g3W21A7SR', 'Nguyễn Văn A', '0919676756', 'admin@gmail.com', 'TPHCM', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('iBqq0XFIR', 'Nguyễn Hoàng Nhân Viên', '0902845188', 'user@gmail.com', 'Thủ Đức', '1de71305d2471d54ec97f4ced8ca46c4', 'TSmnlxOZuhahamTxmvzwIJUVIMZtSDaBOGibbIxDfRpMRLzRCC', 'user', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('NG9h1XKIg', 'Nguyễn Thị Thu Ngân', '0983219471', 'thungan@gmail.com', 'TPHCM', '8905a308326af9c57dc1bc12dfed3d52', 'PXLUpeZBPXdBHHZKbaTTrTNwXAaQYzwZerqBQXpkCXflYvOtEL', 'user', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 0);
INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('R1JpJXKSg', 'Quản Thị Lý', '0902845188', 'manager@gmail.com', 'TPHCM', '793938ddc9eef3bf6620b6c45c344f8a', 'nhFmCdxVksZYDnfyKzarJTxlHZgPcCLoZLvZqkedOiagCwOOmB', 'manager', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 1);

INSERT INTO `Recipe` (`id`) VALUES
('0sPL9lFSgm');
INSERT INTO `Recipe` (`id`) VALUES
('AsEL9lFSRV');
INSERT INTO `Recipe` (`id`) VALUES
('AyPY9_FSgM');
INSERT INTO `Recipe` (`id`) VALUES
('fTz_X_KSgz'),
('hPGPXlKSRz'),
('nGQWClKSRM'),
('nMwZClKSRm'),
('oEJhulFIR'),
('oiNI9lKSgm'),
('rIt9r_KIRm'),
('rItr9lFIgM'),
('TiHSrlFSgV'),
('TmNIr_FIgM'),
('z7qfX_KIg');

INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('0sPL9lFSgm', 'NVLBotSua', 0.002);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('0sPL9lFSgm', 'NVLDuong', 0.2);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('0sPL9lFSgm', 'NVLOLong', 0.02);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('0sPL9lFSgm', 'NVLSua', 0.3),
('AsEL9lFSRV', 'NVLBotSua', 0.002),
('AsEL9lFSRV', 'NVLDuong', 0.3),
('AsEL9lFSRV', 'NVLOLong', 0.035),
('AsEL9lFSRV', 'NVLSua', 0.45),
('AyPY9_FSgM', 'NVLBotSua', 0.002),
('AyPY9_FSgM', 'NVLDuong', 0.2),
('AyPY9_FSgM', 'NVLOLong', 0.025),
('AyPY9_FSgM', 'NVLSua', 0.4),
('fTz_X_KSgz', 'NVLDuong', 0.02),
('fTz_X_KSgz', 'NVLTrung', 1),
('hPGPXlKSRz', 'NVLDen', 0.1),
('nGQWClKSRM', 'NVLNesCafe', 2),
('nGQWClKSRM', 'NVLSua', 0.4),
('nMwZClKSRm', 'NVLNesCafe', 1),
('nMwZClKSRm', 'NVLSua', 0.3),
('oEJhulFIR', 'NVLDao', 1),
('oiNI9lKSgm', 'NVLDao', 2),
('oiNI9lKSgm', 'NVLDuong', 0.2),
('oiNI9lKSgm', 'NVLTuiTraD', 1),
('rIt9r_KIRm', 'NVLNesCafe', 1),
('rItr9lFIgM', 'NVLNesCafe', 2),
('TiHSrlFSgV', 'NVLDao', 4),
('TiHSrlFSgV', 'NVLDuong', 0.4),
('TiHSrlFSgV', 'NVLTuiTraD', 2),
('TmNIr_FIgM', 'NVLDao', 3),
('TmNIr_FIgM', 'NVLDuong', 0.3),
('TmNIr_FIgM', 'NVLTuiTraD', 2),
('z7qfX_KIg', 'NVLTrang', 0.1);

INSERT INTO `Role` (`id`, `name`) VALUES
('admin', 'Chủ quán');
INSERT INTO `Role` (`id`, `name`) VALUES
('user', 'Nhân viên');
INSERT INTO `Role` (`id`, `name`) VALUES
('manager', 'Quản lý');

INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_UP_INFO');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_VIEW');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('manager', 'CAT_VIEW'),
('user', 'CAT_VIEW'),
('admin', 'CUS_CREATE'),
('manager', 'CUS_CREATE'),
('user', 'CUS_CREATE'),
('admin', 'CUS_UP_INFO'),
('manager', 'CUS_UP_INFO'),
('user', 'CUS_UP_INFO'),
('admin', 'CUS_VIEW'),
('manager', 'CUS_VIEW'),
('user', 'CUS_VIEW'),
('admin', 'EXP_CREATE'),
('manager', 'EXP_CREATE'),
('admin', 'EXP_VIEW'),
('manager', 'EXP_VIEW'),
('admin', 'FOD_CREATE'),
('admin', 'FOD_UP_INFO'),
('admin', 'FOD_UP_STATE'),
('manager', 'FOD_UP_STATE'),
('admin', 'FOD_VIEW'),
('manager', 'FOD_VIEW'),
('user', 'FOD_VIEW'),
('admin', 'ICN_CREATE'),
('manager', 'ICN_CREATE'),
('admin', 'ICN_VIEW'),
('manager', 'ICN_VIEW'),
('admin', 'IMP_CREATE'),
('admin', 'IMP_UP_STATE'),
('manager', 'IMP_UP_STATE'),
('admin', 'IMP_VIEW'),
('manager', 'IMP_VIEW'),
('admin', 'ING_CREATE'),
('admin', 'ING_UP'),
('admin', 'ING_VIEW'),
('manager', 'ING_VIEW'),
('admin', 'INV_CREATE'),
('manager', 'INV_CREATE'),
('user', 'INV_CREATE'),
('admin', 'INV_VIEW'),
('manager', 'INV_VIEW'),
('user', 'INV_VIEW'),
('admin', 'RPT_DEBT'),
('manager', 'RPT_DEBT'),
('admin', 'RPT_SALE'),
('manager', 'RPT_SALE'),
('admin', 'RPT_STOCK'),
('manager', 'RPT_STOCK'),
('admin', 'SUP_CREATE'),
('admin', 'SUP_PAY'),
('admin', 'SUP_UP_INFO'),
('admin', 'SUP_VIEW'),
('manager', 'SUP_VIEW'),
('admin', 'TOP_CREATE'),
('admin', 'TOP_UP_INFO'),
('admin', 'TOP_UP_STATE'),
('manager', 'TOP_UP_STATE'),
('admin', 'TOP_VIEW'),
('manager', 'TOP_VIEW'),
('user', 'TOP_VIEW'),
('admin', 'USE_UP_INFO'),
('admin', 'USE_VIEW'),
('manager', 'USE_VIEW');

INSERT INTO `ShopGeneral` (`id`, `name`, `email`, `phone`, `address`, `wifiPass`, `accumulatePointPercent`, `usePointPercent`) VALUES
('shop', 'Coffee shop', '', '', '', 'coffeeshop123', 0.1, 1);


INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('0yEY9_KSR', '0sPY9lKSgZ', 'M', 45000, 50000, 'AyPY9_FSgM');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('0yEY9_KSR', 'AsPYrlFIRz', 'S', 35000, 45000, '0sPL9lFSgm');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('0yEY9_KSR', 'AyEY9_FSR7', 'L', 52000, 55000, 'AsEL9lFSRV');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('9Ip99_FSR', '9Iprr_KSgZ', 'M', 16000, 20000, 'rItr9lFIgM'),
('9Ip99_FSR', '9Strr_KSRz', 'S', 10000, 15000, 'rIt9r_KIRm'),
('nMwZjlKIR', '7MwZjlFIgZ', 'M', 18000, 25000, 'nGQWClKSRM'),
('nMwZjlKIR', 'nGQWjlKSRz', 'S', 12000, 20000, 'nMwZClKSRm'),
('oiNS9_FSg', 'omNS9lFSRZ', 'M', 45000, 50000, 'TmNIr_FIgM'),
('oiNS9_FSg', 'TiHS9_FSR7', 'L', 50000, 55000, 'TiHSrlFSgV'),
('oiNS9_FSg', 'TmHSrlFSRz', 'S', 30000, 45000, 'oiNI9lKSgm');

INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('1cx_BuFIR', 'NVLBo', -1000, 2000, 'Modify', '2023-11-21 05:12:20');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('5NzPEuFIg', 'NVLBotSua', -0.002, 2.992, 'Sell', '2023-12-08 01:12:13');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('5NzPEuFIg', 'NVLDao', -2, 74, 'Sell', '2023-12-08 01:12:13');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('5NzPEuFIg', 'NVLDuong', -0.4, 58.56, 'Sell', '2023-12-08 01:12:13'),
('5NzPEuFIg', 'NVLNesCafe', -1, 94, 'Sell', '2023-12-08 01:12:13'),
('5NzPEuFIg', 'NVLOLong', -0.02, 2.92, 'Sell', '2023-12-08 01:12:13'),
('5NzPEuFIg', 'NVLSua', -0.3, 48.2, 'Sell', '2023-12-08 01:12:13'),
('5NzPEuFIg', 'NVLTuiTraD', -1, 187, 'Sell', '2023-12-08 01:12:13'),
('65uBPuFIg', 'NVLBotSua', -0.002, 2.998, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLDao', -2, 78, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLDuong', -0.4, 59.6, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLNesCafe', -1, 98, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLOLong', -0.02, 2.98, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLSua', -0.3, 49.4, 'Sell', '2023-12-04 15:12:13'),
('65uBPuFIg', 'NVLTuiTraD', -1, 189, 'Sell', '2023-12-04 15:12:13'),
('awrYPXKSg', 'NVLBotSua', -0.002, 2.994, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLDen', -0.1, 99.8, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLDuong', -0.2, 58.96, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLNesCafe', -2, 95, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLOLong', -0.02, 2.94, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLSua', -0.6, 48.5, 'Sell', '2023-12-06 23:12:13'),
('awrYPXKSg', 'NVLTrang', -0.1, 99.7, 'Sell', '2023-12-06 23:12:13'),
('DJNPPXKIg', 'NVLBotSua', -0.002, 2.99, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLDao', -2, 72, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLDen', -0.1, 99.7, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLDuong', -0.42, 58.14, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLOLong', -0.02, 2.9, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLSua', -0.3, 47.9, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLTrang', -0.1, 99.6, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLTrung', -1, 94, 'Sell', '2023-12-09 04:12:13'),
('DJNPPXKIg', 'NVLTuiTraD', -1, 186, 'Sell', '2023-12-09 04:12:13'),
('dsIKBXKIg', 'NVLTuiTraD', -10, 190, 'Export', '2023-11-10 05:03:58'),
('FAoBEuFIg', 'NVLNesCafe', -1, 99, 'Sell', '2023-12-03 05:12:13'),
('FAoBEuFIg', 'NVLSua', -0.3, 49.7, 'Sell', '2023-12-03 05:12:13'),
('flMvfuFIg', 'NVLDao', 100, 100, 'Import', '2023-11-04 23:55:23'),
('flMvfuFIg', 'NVLDen', 100, 100, 'Import', '2023-11-04 23:55:23'),
('flMvfuFIg', 'NVLTrang', 100, 100, 'Import', '2023-11-04 23:55:23'),
('HsJLPXKSg', 'NVLBotSua', -0.002, 2.996, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLDao', -2, 76, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLDen', -0.1, 99.9, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLDuong', -0.44, 59.16, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLNesCafe', -1, 97, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLOLong', -0.02, 2.96, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLSua', -0.3, 49.1, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLTrang', -0.2, 99.8, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLTrung', -2, 95, 'Sell', '2023-12-05 17:12:13'),
('HsJLPXKSg', 'NVLTuiTraD', -1, 188, 'Sell', '2023-12-05 17:12:13'),
('J5CSfXKSR', 'NVLOLong', 3, 3, 'Import', '2023-11-03 16:05:42'),
('J5CSfXKSR', 'NVLSua', 50, 50, 'Import', '2023-11-03 16:05:42'),
('J5CSfXKSR', 'NVLTuiTraD', 200, 200, 'Import', '2023-11-03 16:05:42'),
('kq5VfXKIg', 'NVLBo', 3000, 3000, 'Import', '2023-11-02 22:00:53'),
('kq5VfXKIg', 'NVLBotNang', 4000, 4000, 'Import', '2023-11-02 22:00:53'),
('kq5VfXKIg', 'NVLBotSua', 1, 1, 'Import', '2023-11-02 22:00:53'),
('kq5VfXKIg', 'NVLNesCafe', 100, 100, 'Import', '2023-11-02 22:00:53'),
('ofuFfXKSR', 'NVLDao', -10, 90, 'Export', '2023-11-12 05:04:18'),
('ofuFfXKSR', 'NVLTrung', -3, 97, 'Export', '2023-11-12 05:04:18'),
('RiWlBuKIR', 'NVLBotSua', 1, 2, 'Modify', '2023-11-20 05:04:18'),
('RiWlBuKIR', 'NVLDao', -10, 80, 'Modify', '2023-11-20 05:04:18'),
('uDK_BuKIR', 'NVLBotSua', 1, 3, 'Modify', '2023-11-25 05:12:13'),
('uDK_BuKIR', 'NVLDuong', 10, 60, 'Modify', '2023-11-25 05:12:13'),
('uX3QEXKIR', 'NVLBotSua', -0.002, 2.986, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLDao', -2, 70, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLDen', -0.1, 99.6, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLDuong', -0.42, 57.52, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLOLong', -0.02, 2.86, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLSua', -0.3, 47, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLTrang', -0.1, 99.5, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLTrung', -1, 93, 'Sell', '2023-12-17 04:56:44'),
('uX3QEXKIR', 'NVLTuiTraD', -1, 185, 'Sell', '2023-12-17 04:56:44'),
('wQx_EuFSR', 'NVLNesCafe', -2, 89, 'Sell', '2024-01-10 05:47:19'),
('wQx_EuFSR', 'NVLSua', -0.3, 46.4, 'Sell', '2024-01-10 05:47:19'),
('y9zOBuFSR', 'NVLDuong', 50, 50, 'Import', '2023-11-05 12:12:45'),
('y9zOBuFSR', 'NVLTrung', 100, 100, 'Import', '2023-11-05 12:12:45'),
('YhDlEXFSg', 'NVLBotSua', -0.002, 2.984, 'Sell', '2023-12-30 05:56:44'),
('YhDlEXFSg', 'NVLDuong', -0.2, 57.32, 'Sell', '2023-12-30 05:56:44'),
('YhDlEXFSg', 'NVLNesCafe', -1, 91, 'Sell', '2023-12-30 05:56:44'),
('YhDlEXFSg', 'NVLOLong', -0.02, 2.84, 'Sell', '2023-12-30 05:56:44'),
('YhDlEXFSg', 'NVLSua', -0.3, 46.7, 'Sell', '2023-12-30 05:56:44'),
('Zf8wEXFSg', 'NVLBotSua', -0.002, 2.988, 'Sell', '2023-12-11 04:56:44'),
('Zf8wEXFSg', 'NVLDuong', -0.2, 57.94, 'Sell', '2023-12-11 04:56:44'),
('Zf8wEXFSg', 'NVLNesCafe', -2, 92, 'Sell', '2023-12-11 04:56:44'),
('Zf8wEXFSg', 'NVLOLong', -0.02, 2.88, 'Sell', '2023-12-11 04:56:44'),
('Zf8wEXFSg', 'NVLSua', -0.6, 47.3, 'Sell', '2023-12-11 04:56:44');





INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('NCCAbby', 'Abby', 'abby@gmail.com', '0988923211', 4900000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('NCCBeemart', 'Beemart', 'beemart@gmail.com', '0983219471', 3480000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('NCCHorecavn', 'Horecavn', 'horecavn@gmail.com', '0923253421', 5500000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('NCCTraCS', 'Trà Chính Sơn', 'trachinhson@gmail.com', '0902845188', 0),
('NCCTrumNL', 'Trùm Nguyên Liệu', 'trumnguyenlieu@gmail.com', '0947812391', 1800000);

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('flMvfuFIg', 'NCCHorecavn', 5500000, 5500000, 'Debt', '2023-11-04 23:55:23', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('J5CSfXKSR', 'NCCAbby', 4900000, 4900000, 'Debt', '2023-11-03 16:05:42', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('kq5VfXKIg', 'NCCBeemart', 3480000, 3480000, 'Debt', '2023-11-02 22:00:53', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('y9zOBuFSR', 'NCCTrumNL', 1800000, 1800000, 'Debt', '2023-11-05 12:12:45', 'g3W21A7SR');





INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('2P1hu_KIg', 'Đào Miếng', '', '', 'oEJhulFIR', 1, 5000, 8000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('BTzlu_FSR', 'Pudding trứng', '', '', 'fTz_X_KSgz', 1, 4000, 5000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('hEMEu_KIg', 'Trân châu đen', '', '', 'hPGPXlKSRz', 1, 2500, 5000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('Rnqfu_FSR', 'Trân châu trắng', '', '', 'z7qfX_KIg', 1, 2500, 5000);

INSERT INTO `UnitType` (`id`, `name`, `value`, `measureType`) VALUES
('dv', 'đơn vị', 1, 'Unit');
INSERT INTO `UnitType` (`id`, `name`, `value`, `measureType`) VALUES
('g', 'g', 1, 'Weight');
INSERT INTO `UnitType` (`id`, `name`, `value`, `measureType`) VALUES
('kg', 'kg', 1000, 'Weight');
INSERT INTO `UnitType` (`id`, `name`, `value`, `measureType`) VALUES
('l', 'l', 1000, 'Volume'),
('ml', 'ml', 1, 'Volume');

DELIMITER //

CREATE TRIGGER update_closedAt
BEFORE UPDATE ON ImportNote
FOR EACH ROW
BEGIN
    IF NEW.status != 'InProgress' THEN
        SET NEW.closedAt = CURRENT_TIMESTAMP;
    END IF;
END;

//

DELIMITER ;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;