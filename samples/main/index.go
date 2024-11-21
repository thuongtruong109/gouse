package main

import (
	"flag"

	"github.com/thuongtruong109/gouse/samples"
)

func main() {
	isDevFlag := flag.Bool("isDev", false, "toggle enviroment mode")
	flag.Parse()
	if *isDevFlag {
		arraySample()
		cacheSample()
		chartSample()
		configSample()
		consoleSample()
		cronSample()
		cryptoSample()
		dateSample()
		functionSample()
		helperSample()
		ioSample()
		mathSample()
		netSample()
		numberSample()
		osSample()
		regexSample()
		stringSample()
		structSample()
		typeSample()
	} else {
		print("Please enable isDev flag to run samples")
	}
}

func arraySample() {
	samples.SampleArrayEqual()
	samples.SampleArrayIncludes()
	samples.SampleArrayMost()
	samples.SampleArraySum()
	samples.SampleArrayChunk()
	samples.SampleArrayDiff()
	samples.SampleArrayDrop()
	samples.SampleArrayIndex()
	samples.SampleArrayMerge()
	samples.SampleArrayCompact()

	samples.SampleArrayIntersect()
	samples.SampleArrayMin()
	samples.SampleArrayMax()

	samples.SampleArrayKeyBy()
	samples.SampleArrayIndexBy()
	samples.SampleArrayFilterBy()
	samples.SampleArrayRejectBy()
	samples.SampleArrayFindBy()
	samples.SampleArrayForBy()
	samples.SampleArrayMapBy()
}

func cacheSample() {
	samples.SampleCacheLocal()
	samples.SampleCacheTmp()
}

func chartSample() {
	samples.SampleChartBar()
	samples.SampleChartLine()
	samples.SampleChartPie()
}

func configSample() {
	samples.SampleConfigJson()
	samples.SampleConfigYaml()
	samples.SampleConfigToml()
}

func consoleSample() {
	samples.SampleConsoleCmd()
	samples.SampleConsoleClear()
	samples.SampleConsoleWithColor()
	samples.SampleConsoleBanner()
	samples.SampleConsoleSelect()
	samples.SampleConsoleHelp()

	samples.SampleConsoleList()
	// samples.SampleConsolePaper()
	samples.SampleConsoleProgress()
	samples.SampleConsoleRealtime()
	samples.SampleConsoleChoice()
	samples.SampleConsoleSpinner()
	samples.SampleConsoleSplit()
	samples.SampleConsoleStopwatch()
	samples.SampleConsoleTable()
	// samples.SampleConsoleTab()
	samples.SampleConsoleCountdown()
	samples.SampleConsoleSequence()
	samples.SampleConsoleInline()
	samples.SampleConsoleParallel()
	samples.SampleConsoleDir()
	samples.SampleConsoleGlamour()
}

func cronSample() {
	samples.SampleCronRun()
}

func cryptoSample() {
	samples.SampleCryptoEncode()
	samples.SampleCryptoDecode()
	samples.SampleCryptoEncryptPassword()
	samples.SampleCryptoDecryptPassword()
	samples.SampleCryptoEncryptFile()
	samples.SampleCryptoDecryptFile()
}

func dateSample() {
	samples.SampleDateTime()

	samples.SampleDateISO()
	samples.SampleDateShort()
	samples.SampleDateLong()
	samples.SampleDateUTC()

	samples.SampleDateToSecond()
	samples.SampleDateToMinute()
	samples.SampleDateToHour()
	samples.SampleDateSleepSecond()
	samples.SampleDateSleepMinute()
	samples.SampleDateSleepHour()

	samples.SampleDateClock()
}

func functionSample() {
	samples.SampleFuncDelay()
	samples.SampleFuncRetry()
	samples.SampleFuncRemain()
	samples.SampleFuncInterval()
	samples.SampleFuncLock()
	samples.SampleFuncRunTime()
}

func helperSample() {
	samples.SampleHelperRandomID()
	samples.SampleHelperUUID()

	samples.SampleHelperAutoMdDoc()
}

func ioSample() {
	samples.SampleIoCheckDir()
	samples.SampleIoCreateDir()
	samples.SampleIoRemoveDir()
	samples.SampleIoLsDir()
	samples.SampleIoHierarchyDir()
	samples.SampleIoCurrentDir()

	samples.SampleIoCheckFile()
	samples.SampleIoCreateFile()
	samples.SampleIoRemoveFile()
	samples.SampleIoReadFileByLine()
	samples.SampleIoFileInfo()
	samples.SampleIoRenameFile()
	samples.SampleIoCopyFile()
	samples.SampleIoTruncateFile()
	samples.SampleIoCleanFile()
	samples.SampleIoWriteToFile()
	samples.SampleIoAppendToFile()
	samples.SampleIoFileObj()

	samples.SampleIoCreatePath()
	samples.SampleIoReadPath()
	samples.SampleIoWritePath()

	samples.SampleIoZip()
	samples.SampleIoUnzip()
}

func mathSample() {
	samples.SampleMathIsPrime()
	samples.SampleMathIsEven()
	samples.SampleMathIsOdd()
	samples.SampleMathIsPerfectSquare()

	samples.SampleMathAbs()
	samples.SampleMathFloor()
	samples.SampleMathCeil()
	samples.SampleMathRound()
	samples.SampleMathMin()
	samples.SampleMathMax()
	samples.SampleMathSum()
	samples.SampleMathMean()
	samples.SampleMathOperators()
	samples.SampleMathPower()
	samples.SampleMathFactorial()
	samples.SampleMathRoot()

	samples.SampleMathLog()
	samples.SampleMathPytago()
	samples.SampleMathTrigonometry()
	samples.SampleMathTransition()

	samples.SampleMathRect()
	samples.SampleMathCircle()
	samples.SampleMathTriangle()
	samples.SampleMathSquare()
	samples.SampleMathCube()
	samples.SampleMathSphere()
	samples.SampleMathCylinder()
	samples.SampleMathCone()
	samples.SampleMathTrapezoid()
	samples.SampleMathParallelogram()
	samples.SampleMathRhombus()
	samples.SampleMathEllipse()
	samples.SampleMathPolygon()
}

func mediaSample() {
	samples.SampleMediaCanvas()
	samples.SampleMediaPngToJpg()
}

func netSample() {
	samples.SampleNetOpen()
	samples.SampleNetEncode()
	samples.SampleNetDecode()
	samples.SampleNetCheck()
	samples.SampleNetCheckWithStatusCode()
	samples.SampleNetHeader()
	samples.SampleNetConnectTime()
	samples.SampleNetProxy()
	samples.SampleApiPortScanner()
	samples.SampleApiPortChecker()
}

func numberSample() {
	samples.SampleNumRandom()
	samples.SampleNumClamp()
	samples.SampleNumInRange()
}

func osSample() {
	samples.SampleOsSystem()
	samples.SampleOsDisk()
	samples.SampleOsCpu()
	samples.SampleOsProfile()
}

func regexSample() {
	samples.SampleRegexIsMatch()
	samples.SampleRegexMatch()
	samples.SampleRegexMatchIndex()
}

func stringSample() {
	samples.SampleStringCapitalize()
	samples.SampleStringCamelCase()
	samples.SampleStringSnakeCase()
	samples.SampleStringKebabCase()

	samples.SampleStringIsLetter()
	samples.SampleStringIsDigit()
	samples.SampleStringIncludes()
	samples.SampleStringIsLower()
	samples.SampleStringIsUpper()

	samples.SampleStringSplit()
	samples.SampleStringWords()
	samples.SampleStringReverse()
	samples.SampleStringLower()
	samples.SampleStringUpper()
	samples.SampleStringRepeat()
	samples.SampleStringTruncate()
	samples.SampleStringReplace()
	samples.SampleStringTrim()
	samples.SampleStringTrimBlank()
	samples.SampleStringTrimPrefix()
	samples.SampleStringTrimSuffix()
	samples.SampleStringJoin()
	samples.SampleStringConcat()
	samples.SampleStringSubStr()
	samples.SampleStringSlice()
	samples.SampleStringSplice()
	samples.SampleStringStartsWith()
	samples.SampleStringEndsWith()
	samples.SampleStringEscape()
	samples.SampleStringUnescape()
	samples.SampleStringPad()

	samples.SampleStringCount()
	samples.SampleStringLines()
	samples.SampleStringIndex()
	samples.SampleStringRandom()
	samples.SampleStringAt()
	samples.SampleStringCodePointAt()
	samples.SampleStringCodePoint()
	samples.SampleStringFromCodePointAt()
	samples.SampleStringFromCodePoint()
}

func structSample() {
	samples.SampleStructMerge()
	samples.SampleStructRemove()
	samples.SampleStructAdd()
	samples.SampleStructSet()
	samples.SampleStructGet()

	samples.SampleStructClone()
	samples.SampleStructHas()
}

func typeSample() {
	samples.SampleTypeStringConvert()
	samples.SampleTypeStructConvert()
	samples.SampleTypeCastToString()

	samples.SampleTypeCheck()
	samples.SampleTypeCheckUUID()
	samples.SampleTypeCheckGmail()
	samples.SampleTypeCheckYahoo()
	samples.SampleTypeCheckOutlook()
	samples.SampleTypeCheckEdu()
	samples.SampleTypeCheckEmail()
	samples.SampleTypeCheckUsername()
	samples.SampleTypeCheckPassword()
	samples.SampleTypeCheckPhone()
}
