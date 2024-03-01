# Edgio Go SDK

## Project Standards

![Static Badge](https://img.shields.io/badge/Code_Style-gofmt-blue)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/42dx/edgio-go-sdk?logo=go)
[![semantic-release](https://img.shields.io/badge/Semantic_Release-Conventional_Commits-77f?logo=semantic-release)](https://github.com/semantic-release/semantic-release)
[![conventional-commits](https://img.shields.io/badge/Conventional_Commits-1.0.0-blue.svg?logo=conventionalcommits)](https://conventionalcommits.org)

## Project Meta Data

[![gh-tag](https://img.shields.io/github/v/tag/42dx/edgio-go-sdk?logo=github&label=Latest%20Version&color=orange)](https://github.com/42dx/edgio-go-sdk/releases)
![GitHub Downloads](https://img.shields.io/github/downloads/42dx/edgio-go-sdk/total?logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAACXBIWXMAAAsTAAALEwEAmpwYAAABvUlEQVR4nO3Zu48OURjA4dcqVuJWIpG4lKhcCkJJoiIaGxUtJSqXP4FGQiEioaG00SoEJVFJBKFDIrRIPDKdfI7Z2Z0Z+ebMecqZ5GTeX3Nm5kQURVEUxX+Ftf62JsYC2xMBtsVY4FAiwMEYC5xKBDgZY4FLiQAXYyxwIxHgeowF5hMBHsRY4EUiwPMYC3xOBPgUY4BZ/EoEqK7NRu6w1b9tidzhQE2A/ZE7zNUEOB65w7maAGcjd7haE+BK5A73awLci9zhWU2Ap5E7fKgJ8D5yhhn8qAlQ3ZuJXGGDha2PXGF3gwC7Ilc42iDAkcgVzjQIcDqGDHf1505MO6zC6x6GfzeYswPsWWC7W6yf2BcZ/Pldqssx0JeeR9p7guUxRNiILy2G/4bNMWQ41iLAicgBbi5h+NuRC6zEq0UM/3YwW15T2InvDbe8vZEjnG8Q4EJMExNarrUMD2uGf9x2y+vyeXtZEOvwMTH8V2zqYP3pDlDB4cTx2Fx0YBABKrj2x7K3oiNDCrACL/EGq0cXoIId1ZdjdGhQAfpQAkyYvqI9KwEmTF/RnpUAEzovOjQlQFsGrgQoiqIoiqKItN/2QY+dnO4r2wAAAABJRU5ErkJggg==&color=7777ff&label=Downloads)
[![Repo's Stars](https://img.shields.io/github/stars/42dx/edgio-go-sdk?style=flat&color=yellow&label=Repo%20Stars&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAAACXBIWXMAAAsTAAALEwEAmpwYAAAEJ0lEQVR4nO2ZvW8bdRzGf1Q09r357MT3YikTZSgIwdAWCZA6MYBQ/oIKykCHDlUZujcUlQ2pVMAYJFQQUgckhKCxHb/cxU7sOunZvoiKdkCZUFGGVlS0gd/vQXcXv9upae4SG/mRnsny3ef5+vHd92xCJppoosAEW3sVdf0EGUfB0lTY2kPXlqaScRPq2jxsHag71ubJOAm/Ph+Crf+OuubAAzX1HoqzHBkXwU6cacHvuKp+QMZBAHkGNXWjCV5THXjAUm87r5FRF6r6221Tb8ADluL4LTLqQk1L9YW/pQDr8UUyyoKlvQRLZX3hvQDAuvIKGVXBUr/qqkwXfByozCyQkb1xWcpfPVNvh19zAkw/QiWeIKMmWMrH/SvTBr8243wCji+RURKKsxys+B994RvgLXjg5vQWLE04uOu8reioxY/B0uZgKWdhKdcGVqYd3IP3XJ6+hlLsLMrROazKx1BSdF/uE6grR7Chn4SdOIW6fgE1/Qpq6nVUtQKq6iaq6vbuV5kB8A1wDx4ox4CSZ7YadU1Xotu0KG/SFblAi5HrtBi5goJ0AcuRUyiIJ2HKz+0Ov5H4yF3AWktY9zrQfWMatjKd4H3g2UrDMlhxx4WI5+WGJcAQLw4OYOunUddZE757Heg39eEr0wHeCy+3wHvgJTBTAjUkBlN67wkV0t9FXdvumXoAlWH94Bvg3fCm9A8M8cxw34Oa+iaq2oNRqAxzJy8+RF56Zyj4Zoh1/QSqyr2DrAzz4LdgiG/8J/hWnZQjsOJ3DqIyzIHPi78hKx59KvhmCOf6vB6/5U9l5KEqwzx4G1ludk/wzRC2ImItntyPyjAXXsghFZN9ge94zq3MfBdkZZghOvDfo0i44NaIysynQVSGOfCG+DlADgUC3xGkHLvqZ2WYO3nxauDgzQCl2Cd+VYbtmOaFy/sXoBxb8qMyzHHeM83x6f2BBzmEUvT+XivTghfAcgJoln+Ai/vR/9XpF/dcmXwnPHOc5YGc+ELwAVbk9/2oTBM858G7ATL86cAD0FX5Sz8q44ELTXiW4UGXwl8EH6AYrfhVGbYD7pkDTYdvBgqPLAnTovx44NQ7dhlhC4ZwDlnhHM0JW/0qwxrwS5xrusQ9ds4RXICC9PoT19+8RGle/BoFofmnBkw5RrP8ZzQj/N0JzzXhWdpxGEiFXwsuQFH+cNf11xAzMISXB74/Kx6lGe6nXviw51QYSE6dDywAXZa+7QtuSpsw+d2fVduDpLk5mubudsM7pouhb4ILUJDudj6nin/SvDT/NL1FhRxGeuo8TYfve/AhsGQINDl1JxB4t8dmhDnw1JAoNcUF5Pk9/8aJn/kETU0t0GSIugFuhBh+JDF/qNtPtKgJ1BBvU1PKYTly3PfjJ589Tm9M5eji4V/wA+H9Pv5EE/1f9S9M4cKBWKiMUAAAAABJRU5ErkJggg==)](https://github.com/42dx/edgio-go-sdk)

### [Changelog](CHANGELOG.md)

### [Roadmap](https://github.com/orgs/42dx/projects/1)
