package externalMergeSort

import (
	"fmt"
	"io"
	"os"
	"otus-hw/28.08.2020-linear-sort/readerFile"
	"otus-hw/28.08.2020-linear-sort/simpleMergeSort"
	"path/filepath"
)

type PartManager interface {
	SplitFile(file io.Reader)
	createFile()
	SortParts()
	createPart(file io.Writer)
	MergeParts(first io.Writer, second io.Writer)
}

type partManager struct {
	index     uint64
	fileNames []string
	fileName  string
}

func (p *partManager) IncrementIndex() {
	p.index++
}
func (p *partManager) GetIndex() uint64 {
	return p.index
}
func MergeSort(file io.Reader) (err error) {
	p := new(partManager)
	p.fileName = "C:/Users/pogos/workspace/go/otus/otus-hw/28.08.2020-linear-sort/externalMergeSort/tmpFiles/"
	RemoveContents(p.fileName)
	err = p.SplitFile(file, 10000)
	if err != nil {
		return err
	}
	//sort all parts using simple mergeSort
	p.SortParts()
	p.MergeSort()

	return nil
}
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *partManager) MergeSort() (err error) {
	for {
		if len(p.fileNames) < 2 {
			return
		}
		firstReader, _ := os.Open(p.fileNames[0] + ".bin")
		secondReader, _ := os.Open(p.fileNames[1] + ".bin")
		p.IncrementIndex()
		newFileName := p.fileName + fmt.Sprint(p.GetIndex())
		newFile, _ := os.Create(p.fileName + fmt.Sprint(p.GetIndex()) + ".bin")
		newFileTxt, _ := os.Create(p.fileName + fmt.Sprint(p.GetIndex()) + ".txt")
		p.MergeParts(firstReader, secondReader, newFile, newFileTxt)
		firstReader.Close()
		secondReader.Close()
		newFile.Close()
		newFileTxt.Close()
		os.Remove(p.fileNames[0] + ".bin")
		os.Remove(p.fileNames[1] + ".bin")
		os.Remove(p.fileNames[0] + ".txt")
		os.Remove(p.fileNames[1] + ".txt")
		p.fileNames = p.fileNames[2:]
		p.fileNames = append(p.fileNames, newFileName)
		fmt.Println("done   ", p.GetIndex())
	}
}
func (p *partManager) SplitFile(file io.Reader, numCount int) (err error) {
	var num uint16
	bts := make([]byte, 2)
	nums := []uint16{}
	for {
		num, _, err = readerFile.ReadUint16(file, bts)
		if err != nil {
			if err == io.EOF {
				err := p.createBinAndTxtParts(nums)
				if err != nil {
					return err
				}
				break
			}
			return err
		}
		nums = append(nums, num)
		if len(nums) >= numCount {
			err := p.createBinAndTxtParts(nums)
			if err != nil {
				return err
			}
			p.IncrementIndex()
			nums = []uint16{}
		}
	}
	return nil
}

func (p *partManager) createBinAndTxtParts(nums []uint16) (err error) {

	fileName := p.fileName + fmt.Sprint(p.GetIndex())
	newFile, err := os.Create(fileName + ".bin")
	err = p.createPart(newFile, nums)
	if err != nil {
		return err
	}
	defer newFile.Close()
	newFileTxt, err := os.Create(fileName + ".txt")
	err = p.createPartTxt(newFileTxt, nums)
	if err != nil {
		return err
	}
	defer newFileTxt.Close()
	p.fileNames = append(p.fileNames, fileName)

	return nil
}

func (p *partManager) createPart(writer io.Writer, nums []uint16) (err error) {

	bts := make([]byte, 2)
	for _, num := range nums {
		_, err = readerFile.WriteUint16(writer, num, bts)
		if err != nil {
			return err
		}
	}

	return nil
}
func (p *partManager) createPartTxt(writer io.Writer, nums []uint16) (err error) {
	for _, num := range nums {
		if _, err = fmt.Fprint(writer, num, " "); err != nil {
			return err
		}
	}

	return nil
}

func (p *partManager) SortParts() (err error) {
	uints := []uint16{}
	sortedUints := []uint16{}
	for _, fileName := range p.fileNames {
		file, err := os.Open(fileName + ".bin")
		if err != nil {
			return err
		}
		uints, _, err = readerFile.ReadAllUint16(file)
		if err != nil {
			return err
		}
		sortedUints = simpleMergeSort.MergeSort(uints)

		file.Close()

		newFile, err := os.Create(fileName + ".bin")
		err = p.createPart(newFile, sortedUints)
		if err != nil {
			return err
		}
		newFile.Close()
		newFileTxt, err := os.Create(fileName + ".txt")
		err = p.createPartTxt(newFileTxt, sortedUints)
		if err != nil {
			return err
		}
		newFileTxt.Close()
	}
	return nil
}
func (p *partManager) MergeParts(leftReader io.Reader, rightReader io.Reader, res io.Writer, resTxt io.Writer) (err error) {

	buf := make([]byte, 2)
	lNum := uint16(0)
	rNum := uint16(0)
	isLeftEnd := false
	isRightEnd := false

	lNum, _, err = readerFile.ReadUint16(leftReader, buf)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	rNum, _, err = readerFile.ReadUint16(rightReader, buf)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	for {
		if lNum < rNum {
			readerFile.WriteUint16(res, lNum, buf)
			if _, err = fmt.Fprint(resTxt, lNum, " "); err != nil {
				return err
			}
			lNum, _, err = readerFile.ReadUint16(leftReader, buf)
			if err != nil {
				if err == io.EOF {
					isLeftEnd = true
					break
				} else {
					return err
				}
			}
		} else {
			readerFile.WriteUint16(res, rNum, buf)
			if _, err = fmt.Fprint(resTxt, rNum, " "); err != nil {
				return err
			}
			rNum, _, err = readerFile.ReadUint16(rightReader, buf)
			if err != nil {
				if err == io.EOF {
					isRightEnd = true
					break
				} else {
					return err
				}
			}
		}
	}
	if isLeftEnd {
		for {
			readerFile.WriteUint16(res, rNum, buf)
			if _, err = fmt.Fprint(resTxt, rNum, " "); err != nil {
				return err
			}
			rNum, _, err = readerFile.ReadUint16(rightReader, buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
		}
	}

	if isRightEnd {
		for {
			readerFile.WriteUint16(res, lNum, buf)
			if _, err = fmt.Fprint(resTxt, lNum, " "); err != nil {
				return err
			}
			lNum, _, err = readerFile.ReadUint16(leftReader, buf)
			if err != nil {
				if err == io.EOF {
					break
				} else {
					return err
				}
			}
		}
	}
	return
}
