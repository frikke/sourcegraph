import React, { useState, useCallback } from 'react'

import ChevronDoubleLeftIcon from 'mdi-react/ChevronDoubleLeftIcon'
import ChevronDoubleRightIcon from 'mdi-react/ChevronDoubleRightIcon'
import { useLocation } from 'react-router'
import { animated, useSpring } from 'react-spring'

import { Button, useLocalStorage, Icon, Link, Text } from '@sourcegraph/wildcard'

import { Scalars } from '../../../../../graphql-operations'
import { insertNameIntoLibraryItem, insertQueryIntoLibraryItem } from '../../yaml-util'

import combySample from './comby.batch.yaml'
import goImportsSample from './go-imports.batch.yaml'
import helloWorldSample from './hello-world.batch.yaml'
import minimalSample from './minimal.batch.yaml'
import { ReplaceSpecModal } from './ReplaceSpecModal'

import styles from './LibraryPane.module.scss'

interface LibraryItem {
    name: string
    code: string
}

const LIBRARY: [LibraryItem, LibraryItem, LibraryItem, LibraryItem] = [
    { name: 'hello world', code: helloWorldSample },
    { name: 'minimal', code: minimalSample },
    { name: 'modify with comby', code: combySample },
    { name: 'update go imports', code: goImportsSample },
]

const LIBRARY_PANE_DEFAULT_COLLAPSED = 'batch-changes.ssbc-library-pane-default-collapsed'
// Match to `.collapse-button` class width
const BUTTON_WIDTH = '1.25rem'
// Match to `.list-container` class width
const CONTENT_WIDTH = '14rem'

type LibraryPaneProps =
    | {
          /**
           * The name of the batch change, used for automatically filling in the name for any
           * item selected from the library.
           */
          name: Scalars['String']
          onReplaceItem: (item: string) => void
          isReadOnly?: false
      }
    | {
          name: Scalars['String']
          isReadOnly: true
      }

export const LibraryPane: React.FunctionComponent<React.PropsWithChildren<LibraryPaneProps>> = ({ name, ...props }) => {
    // Remember the last collapsed state of the pane
    const [defaultCollapsed, setDefaultCollapsed] = useLocalStorage(LIBRARY_PANE_DEFAULT_COLLAPSED, false)
    const [collapsed, setCollapsed] = useState(defaultCollapsed)
    const [selectedItem, setSelectedItem] = useState<LibraryItem>()

    const [containerStyle, animateContainer] = useSpring(() => ({
        width: collapsed ? BUTTON_WIDTH : CONTENT_WIDTH,
    }))

    const [headerStyle, animateHeader] = useSpring(() => ({
        opacity: collapsed ? 0 : 1,
        width: collapsed ? '0rem' : CONTENT_WIDTH,
    }))

    const [contentStyle, animateContent] = useSpring(() => ({
        display: collapsed ? 'none' : 'block',
        opacity: collapsed ? 0 : 1,
    }))

    const toggleCollapse = useCallback(
        (collapsed: boolean) => {
            setCollapsed(collapsed)
            setDefaultCollapsed(collapsed)
            animateContainer.start({ width: collapsed ? BUTTON_WIDTH : CONTENT_WIDTH })
            animateContent.start({
                /* eslint-disable callback-return */
                // We need the display: none property change to happen in sequence *after*
                // the opacity property change or else the content will disappear
                // immediately. This use of the API is following the suggestion from
                // https://react-spring.io/hooks/use-spring#this-is-how-you-create-a-script
                to: async next => {
                    await next({ display: 'block', opacity: collapsed ? 0 : 1 })
                    if (collapsed) {
                        await next({ display: 'none' })
                    }
                },
                /* eslint-enable callback-return */
            })
            animateHeader.start({ opacity: collapsed ? 0 : 1, width: collapsed ? '0rem' : CONTENT_WIDTH })
        },
        [animateContainer, animateContent, animateHeader, setDefaultCollapsed]
    )

    const location = useLocation()
    const parameters = new URLSearchParams(location.search)
    const searchQuery = parameters.get('q') ?? undefined

    const updateLibraryItemFields = useCallback(
        (code: string, name: string): string => {
            const updatedCode = insertNameIntoLibraryItem(code, name)

            if (searchQuery) {
                return insertQueryIntoLibraryItem(updatedCode, searchQuery)
            }

            return updatedCode
        },
        [searchQuery]
    )

    const onConfirm = useCallback(() => {
        if (selectedItem && !('isReadOnly' in props && props.isReadOnly)) {
            const updatedCode = updateLibraryItemFields(selectedItem.code, name)
            props.onReplaceItem(updatedCode)
            setSelectedItem(undefined)
        }
    }, [name, selectedItem, props, updateLibraryItemFields])

    return (
        <>
            {selectedItem ? (
                <ReplaceSpecModal
                    libraryItemName={selectedItem.name}
                    onCancel={() => setSelectedItem(undefined)}
                    onConfirm={onConfirm}
                />
            ) : null}
            <animated.div style={containerStyle} className="d-flex flex-column mr-3">
                <div className={styles.header}>
                    <animated.h4 className="m-0" style={headerStyle}>
                        Library
                    </animated.h4>
                    <div className={styles.collapseButton}>
                        <Button
                            className="p-0"
                            onClick={() => toggleCollapse(!collapsed)}
                            aria-label={collapsed ? 'Expand' : 'Collapse'}
                        >
                            <Icon
                                role="img"
                                aria-hidden={true}
                                as={collapsed ? ChevronDoubleRightIcon : ChevronDoubleLeftIcon}
                            />
                        </Button>
                    </div>
                </div>

                <animated.div style={contentStyle}>
                    <ul className={styles.listContainer}>
                        {LIBRARY.map(item => (
                            <li className={styles.libraryItem} key={item.name}>
                                <Button
                                    className={styles.libraryItemButton}
                                    disabled={'isReadOnly' in props && props.isReadOnly}
                                    onClick={() => setSelectedItem(item)}
                                >
                                    {item.name}
                                </Button>
                            </li>
                        ))}
                    </ul>
                    <Text className={styles.lastItem}>
                        <Link to="https://github.com/sourcegraph/batch-change-examples">View more examples</Link>
                    </Text>
                </animated.div>
            </animated.div>
        </>
    )
}
