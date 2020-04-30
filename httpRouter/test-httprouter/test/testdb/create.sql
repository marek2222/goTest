USE [master]
GO
CREATE DATABASE [SamolotyMiejscaTest]
GO
ALTER DATABASE [SamolotyMiejscaTest] SET RECOVERY SIMPLE 
GO

GO
USE SamolotyMiejscaTest
GO
CREATE SCHEMA miejsca
GO
CREATE TABLE miejsca.Samoloty(
	SamolotID int NOT NULL,
	MiejscaJson nvarchar(max) NULL,
	CONSTRAINT [PK_Samoloty] PRIMARY KEY CLUSTERED (SamolotID ASC)
)
GO
INSERT INTO miejsca.Samoloty (SamolotID,MiejscaJson) VALUES (123,null)
INSERT INTO miejsca.Samoloty (SamolotID,MiejscaJson) VALUES (456,null)
INSERT INTO miejsca.Samoloty (SamolotID,MiejscaJson) VALUES (789,null)
GO