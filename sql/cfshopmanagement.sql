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
  `price` float NOT NULL,
  `amountImport` float NOT NULL,
  `totalUnit` int NOT NULL,
  `unitTypeName` text NOT NULL,
  PRIMARY KEY (`importNoteId`,`ingredientId`),
  KEY `ingredientId` (`ingredientId`),
  CONSTRAINT `ImportNoteDetail_ibfk_1` FOREIGN KEY (`ingredientId`) REFERENCES `Ingredient` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `Ingredient`;
CREATE TABLE `Ingredient` (
  `id` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `name` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `amount` int DEFAULT '0',
  `unitTypeId` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` float NOT NULL,
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
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS `StockReportDetail`;
CREATE TABLE `StockReportDetail` (
  `reportId` varchar(12) NOT NULL,
  `ingredientId` varchar(12) NOT NULL,
  `initial` int NOT NULL,
  `sell` int NOT NULL,
  `import` int NOT NULL,
  `export` int NOT NULL,
  `modify` int NOT NULL,
  `final` int NOT NULL,
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
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
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
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
  `createdAt` datetime DEFAULT CURRENT_TIMESTAMP,
  `updatedAt` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deletedAt` datetime DEFAULT NULL,
  `isActive` tinyint(1) DEFAULT '1',
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
('kShzQ1FIR', 'Nếu không đổi tên thì xóa trường này', 'Nếu ko đổi mô tả thì xóa trường này', 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('LNQxrPFSg', '21521495', '', 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('nk7XQ1KIR', 'Tên danh mục', 'Mô tả danh mục', 0);
INSERT INTO `Category` (`id`, `name`, `description`, `amountProduct`) VALUES
('QJ8y9EFIR', 'cà phê', '', 1),
('XTWY9PKSR', 'sữa', '', 0);

INSERT INTO `CategoryFood` (`foodId`, `categoryId`) VALUES
('_3My6PKIR', 'QJ8y9EFIR');


INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('_1V3zQFSR', 'quao', 'a@gmail.com', '0987687114', 10);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('4SlRmwFSg', 'hi', 'b@gmail.com', '0398750911', 290);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('IdCustomer', 'Nếu không sửa tên thì xóa trường này', 'a@gmail.comm', '0987658712', 0);
INSERT INTO `Customer` (`id`, `name`, `email`, `phone`, `point`) VALUES
('KH001', 'tên khách hàng', 'a@gmail.com', '01234567893', 0),
('l-ZYiwFSR', 'hi', 'a@gmail.com', '0901982312', 0);

INSERT INTO `ExportNote` (`id`, `reason`, `createdBy`, `createdAt`) VALUES
('UEl6HYFIg', 'Damaged', 'g3W21A7SR', '2024-01-07 17:50:03');


INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`, `unitTypeName`) VALUES
('UEl6HYFIg', '0_LcnYFIR', 2, 'l');
INSERT INTO `ExportNoteDetail` (`exportNoteId`, `ingredientId`, `amountExport`, `unitTypeName`) VALUES
('UEl6HYFIg', '8Se57YKSR', 1, 'đơn vị');


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
('FOD_CREATE', 'Tạo sản phẩm', 'Sản phẩm'),
('FOD_UP_INFO', 'Chỉnh sửa thông tin sản phẩm', 'Sản phẩm'),
('FOD_UP_STATE', 'Chỉnh sửa trạng thái sản phẩm', 'Sản phẩm'),
('FOD_VIEW', 'Xem sản phẩm', 'Sản phẩm'),
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
('USE_UP_INFO', 'Chỉnh sửa thông tin người dùng', 'Người dùng'),
('USE_UP_STATE', 'Chỉnh sửa trạng thái người dùng', 'Người dùng'),
('USE_VIEW', 'Xem người dùng', 'Người dùng');

INSERT INTO `Food` (`id`, `name`, `description`, `cookingGuide`, `image`, `isActive`) VALUES
('_3My6PKIR', 'Cà phê sữa đá', 'Cà phê pha với sữa', 'none', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Food%2Fno_orderW_ZseEFIR?alt=media', 1);


INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('-fRh7LFIR', 'SupCake0001', 50000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2024-01-07 16:47:54', '2024-01-07 17:30:20');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('-sRhnYKIg', 'SupCake0001', 50000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2024-01-07 16:47:54', '2024-01-07 17:30:24');
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('8xZUtEFIg', 'SupCake0001', 420, 'InProgress', 'g3W21A7SR', NULL, '2024-01-08 04:32:27', NULL);
INSERT INTO `ImportNote` (`id`, `supplierId`, `totalPrice`, `status`, `createdBy`, `closedBy`, `createdAt`, `closedAt`) VALUES
('uI3h7LFSg', 'SupCoffe0001', 3000, 'Done', 'g3W21A7SR', 'g3W21A7SR', '2024-01-07 16:48:24', '2024-01-07 17:30:14');

INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('-fRh7LFIR', '0_LcnYFIR', 20000, 2.5, 50000, 'l');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('-sRhnYKIg', '0_LcnYFIR', 20000, 2.5, 50000, 'l');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('8xZUtEFIg', '0_LcnYFIR', 21, 20, 420, 'ml');
INSERT INTO `ImportNoteDetail` (`importNoteId`, `ingredientId`, `price`, `amountImport`, `totalUnit`, `unitTypeName`) VALUES
('uI3h7LFSg', '8Se57YKSR', 2500, 1.2, 3000, 'đơn vị');

INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('0_LcnYFIR', 'Sữa', 88, 'l', 0);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('8Se57YKSR', 'Trứng', 27, 'dv', 2500);
INSERT INTO `Ingredient` (`id`, `name`, `amount`, `unitTypeId`, `price`) VALUES
('Ecd5nLKSR', 'Đường', 28, 'kg', 35000);









INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('g3W21A7SR', 'Nguyễn Văn A', '0919676756', 'admin@gmail.com', 'TPHCM', '5e107317df151f6e8e0015c4f2ee7936', 'mVMxRDAHpAJfyzuiXWRELghNpynUqBKueSboGBcrwHUuzEWsms', 'admin', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 1);
INSERT INTO `MUser` (`id`, `name`, `phone`, `email`, `address`, `password`, `salt`, `roleId`, `image`, `isActive`) VALUES
('za1u8m4Sg', 'Nguyễn Văn U', '0966656041', 'user@gmail.com', 'TPHCM', 'cb58ac5a2272517d1960565444bde187', 'QYlnGKRgYBxIXzMnnQSVcglbtjPsAhVlxMRMDaqnaquxwADSur', 'user', 'https://firebasestorage.googleapis.com/v0/b/coffee-shop-web.appspot.com/o/Default%2Favatar.jpg?alt=media', 1);


INSERT INTO `Recipe` (`id`) VALUES
('DUS0dyFIgz');
INSERT INTO `Recipe` (`id`) VALUES
('Gg-fvsKIRz');
INSERT INTO `Recipe` (`id`) VALUES
('l3GseEFIgm');
INSERT INTO `Recipe` (`id`) VALUES
('l3Ms6EKIRM'),
('lHc0IyKIgz'),
('lrvAvsKSgz');

INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('DUS0dyFIgz', '0_LcnYFIR', 0.9);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('DUS0dyFIgz', '8Se57YKSR', 1);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('DUS0dyFIgz', 'Ecd5nLKSR', 0.5);
INSERT INTO `RecipeDetail` (`recipeId`, `ingredientId`, `amountNeed`) VALUES
('Gg-fvsKIRz', '0_LcnYFIR', 1.2),
('Gg-fvsKIRz', 'Ecd5nLKSR', 3.4),
('l3GseEFIgm', '0_LcnYFIR', 0.3),
('l3Ms6EKIRM', '0_LcnYFIR', 0.5),
('lHc0IyKIgz', '8Se57YKSR', 0.012),
('lHc0IyKIgz', 'Ecd5nLKSR', 0.006),
('lrvAvsKSgz', '0_LcnYFIR', 0.1),
('lrvAvsKSgz', '8Se57YKSR', 0.2),
('lrvAvsKSgz', 'Ecd5nLKSR', 0.3);

INSERT INTO `Role` (`id`, `name`) VALUES
('admin', 'admin');
INSERT INTO `Role` (`id`, `name`) VALUES
('user', 'user');


INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_CREATE');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_UP_INFO');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CAT_VIEW');
INSERT INTO `RoleFeature` (`roleId`, `featureId`) VALUES
('admin', 'CUS_CREATE'),
('admin', 'CUS_UP_INFO'),
('admin', 'CUS_VIEW'),
('admin', 'EXP_CREATE'),
('admin', 'EXP_VIEW'),
('admin', 'FOD_CREATE'),
('admin', 'FOD_UP_INFO'),
('admin', 'FOD_UP_STATE'),
('admin', 'FOD_VIEW'),
('admin', 'ICN_CREATE'),
('admin', 'ICN_VIEW'),
('admin', 'IMP_CREATE'),
('admin', 'IMP_UP_STATE'),
('admin', 'IMP_VIEW'),
('admin', 'ING_CREATE'),
('admin', 'ING_UP'),
('admin', 'ING_VIEW'),
('admin', 'INV_CREATE'),
('admin', 'INV_VIEW'),
('admin', 'RPT_DEBT'),
('admin', 'RPT_SALE'),
('admin', 'RPT_STOCK'),
('admin', 'SUP_CREATE'),
('admin', 'SUP_PAY'),
('admin', 'SUP_UP_INFO'),
('admin', 'SUP_VIEW'),
('admin', 'TOP_CREATE'),
('admin', 'TOP_UP_INFO'),
('admin', 'TOP_UP_STATE'),
('admin', 'TOP_VIEW'),
('admin', 'USE_UP_INFO'),
('admin', 'USE_UP_STATE'),
('admin', 'USE_VIEW');

INSERT INTO `ShopGeneral` (`id`, `name`, `email`, `phone`, `address`, `wifiPass`, `accumulatePointPercent`, `usePointPercent`) VALUES
('shop', 'Coffee shop', '', '', '', 'coffeeshop123', 0.001, 1);


INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('_3My6PKIR', 'l3Gs6EKSRz', 'S', 7000, 19000, 'l3GseEFIgm');
INSERT INTO `SizeFood` (`foodId`, `sizeId`, `name`, `cost`, `price`, `recipeId`) VALUES
('_3My6PKIR', 'l3MseEKSgZ', 'M', 8000, 24000, 'l3Ms6EKIRM');


INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('-fRh7LFIR', '0_LcnYFIR', 2.5, 2.5, 'Import', '2024-01-07 17:30:20');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('-sRhnYKIg', '0_LcnYFIR', 2.5, 4.5, 'Import', '2024-01-07 17:30:24');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('CO3JNQKSR', '0_LcnYFIR', -20, 19.7, 'Sell', '2024-01-09 06:59:26');
INSERT INTO `StockChangeHistory` (`id`, `ingredientId`, `amount`, `amountLeft`, `type`, `createdAt`) VALUES
('CoyrdQKIR', '0_LcnYFIR', -100, 96.2, 'Sell', '2024-01-09 07:40:16'),
('CoyrdQKIR', '8Se57YKSR', -33, 30.988, 'Sell', '2024-01-09 07:40:16'),
('CoyrdQKIR', 'Ecd5nLKSR', -40, 35.594, 'Sell', '2024-01-09 07:40:16'),
('Ekn6vQKSR', '0_LcnYFIR', -15, 12.4, 'Sell', '2024-01-09 07:24:07'),
('Ekn6vQKSR', '8Se57YKSR', -38, 37, 'Sell', '2024-01-09 07:24:07'),
('Ekn6vQKSR', 'Ecd5nLKSR', -22, 18.1, 'Sell', '2024-01-09 07:24:07'),
('FqkUOQKSR', '0_LcnYFIR', -3, 0.0999999, 'Sell', '2024-01-09 07:37:43'),
('FqkUOQKSR', '8Se57YKSR', -34, 33, 'Sell', '2024-01-09 07:37:43'),
('FqkUOQKSR', 'Ecd5nLKSR', -6, 2.1, 'Sell', '2024-01-09 07:37:43'),
('hHQWdwFIR', '0_LcnYFIR', -12, 9.4, 'Sell', '2024-01-09 07:26:39'),
('hHQWdwFIR', '8Se57YKSR', -37, 36, 'Sell', '2024-01-09 07:26:39'),
('hHQWdwFIR', 'Ecd5nLKSR', -18, 14.1, 'Sell', '2024-01-09 07:26:39'),
('j24sOQKSR', '0_LcnYFIR', -9, 6.4, 'Sell', '2024-01-09 07:37:15'),
('j24sOQKSR', '8Se57YKSR', -36, 35, 'Sell', '2024-01-09 07:37:15'),
('j24sOQKSR', 'Ecd5nLKSR', -14, 10.1, 'Sell', '2024-01-09 07:37:15'),
('JmQjdQKSg', '0_LcnYFIR', -96, 92.2, 'Sell', '2024-01-09 07:40:51'),
('JmQjdQKSg', '8Se57YKSR', -31, 28.988, 'Sell', '2024-01-09 07:40:51'),
('JmQjdQKSg', 'Ecd5nLKSR', -36, 31.594, 'Sell', '2024-01-09 07:40:51'),
('NJUydwFIR', '0_LcnYFIR', -6, 3.1, 'Sell', '2024-01-09 07:37:34'),
('NJUydwFIR', '8Se57YKSR', -35, 34, 'Sell', '2024-01-09 07:37:34'),
('NJUydwFIR', 'Ecd5nLKSR', -10, 6.1, 'Sell', '2024-01-09 07:37:34'),
('P-6cKQKSR', '0_LcnYFIR', -92, 88.2, 'Sell', '2024-01-09 07:49:09'),
('P-6cKQKSR', '8Se57YKSR', -29, 26.988, 'Sell', '2024-01-09 07:49:09'),
('P-6cKQKSR', 'Ecd5nLKSR', -32, 27.594, 'Sell', '2024-01-09 07:49:09'),
('qRgaHwFIR', '0_LcnYFIR', -20, 17.3, 'Sell', '2024-01-09 07:00:00'),
('qRgaHwFIR', '8Se57YKSR', -40, 39, 'Sell', '2024-01-09 07:00:00'),
('qRgaHwFIR', 'Ecd5nLKSR', -30, 26.1, 'Sell', '2024-01-09 07:00:00'),
('UEl6HYFIg', '0_LcnYFIR', -2, 2, 'Export', '2024-01-07 17:50:03'),
('UEl6HYFIg', '8Se57YKSR', -1, 0, 'Export', '2024-01-07 17:50:03'),
('uI3h7LFSg', '8Se57YKSR', 1.2, 1.2, 'Import', '2024-01-07 17:30:14'),
('WYQjDwFSR', '0_LcnYFIR', -17, 14.6, 'Sell', '2024-01-09 07:23:23'),
('WYQjDwFSR', '8Se57YKSR', -39, 38, 'Sell', '2024-01-09 07:23:23'),
('WYQjDwFSR', 'Ecd5nLKSR', -26, 22.1, 'Sell', '2024-01-09 07:23:23');

INSERT INTO `StockReport` (`id`, `timeFrom`, `timeTo`, `createdAt`, `updatedAt`, `deletedAt`, `isActive`) VALUES
('LRYQ5EFSR', '2023-11-30 17:00:00', '2023-12-31 16:59:59', '2024-01-08 04:15:49', '2024-01-08 04:15:49', NULL, 1);




INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCacao0001', 'NCC Cacao', 'cacao@gmail.com', '0905555555', -1248000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCake0001', 'NCC Bánh', 'banh@gmail.com', '0943334445', -1935000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupCoffe0001', 'NCC Cà Phê', 'caphe@gmail.com', '0901234567', -77000);
INSERT INTO `Supplier` (`id`, `name`, `email`, `phone`, `debt`) VALUES
('SupHoney0001', 'NCC Mật Ong', 'matong@gmail.com', '0927777777', -2000),
('SupIceCr0001', 'NCC Kem', 'kem@gmail.com', '0999999999', -60000),
('SupMilk0001', 'NCC Sữa Tuyệt trùng', 'suatuyettrung@gmail.com', '0919876543', -325000),
('SupOTea0001', 'NCC Trà Ôlong', 'olong@gmail.com', '0922333445', -30000),
('SupPearl0001', 'NCC Trân Châu', 'tranchau@gmail.com', '0911122334', -3500),
('SupSugar0001', 'NCC Đường', 'duong@gmail.com', '0921112223', -30000),
('SupTea0001', 'NCC Trà', 'tra@gmail.com', '0922233445', -383000);

INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('-fRh7LFIR', 'SupCake0001', 50000, -1985000, 'Debt', '2024-01-07 17:30:20', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('-sRhnYKIg', 'SupCake0001', 50000, -1935000, 'Debt', '2024-01-07 17:30:24', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN001', 'SupCake0001', -2035000, -2035000, 'Debt', '2023-12-26 05:16:43', 'g3W21A7SR');
INSERT INTO `SupplierDebt` (`id`, `supplierId`, `amount`, `amountLeft`, `type`, `createdAt`, `createdBy`) VALUES
('PN006', 'SupCacao0001', -1198000, -1248000, 'Debt', '2023-12-26 05:17:40', 'g3W21A7SR'),
('PN010', 'SupTea0001', -363000, -383000, 'Debt', '2023-12-26 05:17:23', 'g3W21A7SR'),
('PN011', 'SupMilk0001', -320000, -325000, 'Debt', '2023-12-26 05:15:59', 'g3W21A7SR'),
('uI3h7LFSg', 'SupCoffe0001', 3000, -77000, 'Debt', '2024-01-07 17:30:14', 'g3W21A7SR');





INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('_rDADyFSR', 'Trân châu trắng', 'Topping trân châu trắng', 'nấu đi nào', 'lrvAvsKSgz', 1, 1000, 5000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('lHc0IyKSg', 'Trân châu', '', '', 'lHc0IyKIgz', 1, 1000, 4000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('Mg-BvsKIR', 'Hạt sen', 'hạt sen', 'nấu hạt sen', 'Gg-fvsKIRz', 1, 3000, 6000);
INSERT INTO `Topping` (`id`, `name`, `description`, `cookingGuide`, `recipeId`, `isActive`, `cost`, `price`) VALUES
('vUSAdsFIR', 'Pudding', 'Mềm', 'Đánh trứng', 'DUS0dyFIgz', 1, 4000, 10000);

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